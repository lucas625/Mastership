<template>
  <v-container
    class="fill-height"
    fluid
    style="background-color: #121212;"
  >
    <alert
        :active="alertActive"
        :message="alertMessage"
        :message-type="alertMessageType"
        @close="clearAlert"
        class="alert-class"
    />
    <loading-bar
        :loading-message="loadingMessage"
        :is-loading-props="isLoading"
    />
    <v-row justify="center">
      <v-col align="center">
        <v-card max-width="600px">
          <v-card-title class="mb-5">
            <v-row justify="center">
              <v-col align="center">
                <h1>
                  {{ hasResults ? "Analyze the results" : "Run the experiment" }}
                </h1>
              </v-col>
            </v-row>
          </v-card-title>
          <v-card-text>
            <v-row justify="center">
              <v-col align="center">
                <v-form ref="form" v-model="isFormValid">
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="label"
                          label="Label"
                          :rules="[FORM_RULES.ruleRequiredField]"
                      />
                      <v-text-field
                          v-model="interactions"
                          label="Number of requests"
                          :rules="[FORM_RULES.ruleRequiredField, FORM_RULES.ruleMinMaxValueField(100, 100000)]"
                          type="Number"
                          :disabled="hasResults"
                      />
                      <v-text-field
                          v-model="batchSize"
                          label="Size of groups of requests"
                          :rules="[FORM_RULES.ruleRequiredField, FORM_RULES.ruleMinMaxValueField(1, 10000)]"
                          type="Number"
                          :disabled="hasResults"
                      />
                      <v-text-field
                          v-model="intervalBetweenBatchesInMilliseconds"
                          label="Interval in ms between groups of requests"
                          :rules="[FORM_RULES.ruleMinMaxValueField(0, 1000)]"
                          type="Number"
                          :disabled="hasResults"
                      />
                      <v-text-field
                          v-if="hasResults"
                          v-model="mean"
                          label="Mean in Microseconds"
                          :rules="[FORM_RULES.ruleRequiredField]"
                          type="Number"
                          :disabled="true"
                      />
                      <v-text-field
                          v-if="hasResults"
                          v-model="standardDeviation"
                          label="Standard deviation in Microseconds"
                          :rules="[FORM_RULES.ruleRequiredField]"
                          type="Number"
                          :disabled="true"
                      />
                      <v-text-field
                          v-if="hasResults"
                          v-model="failures"
                          label="Total number of request failures on test"
                          type="Number"
                          :disabled="true"
                      />
                      <v-textarea
                          v-if="hasResults"
                          v-model="rttsInMicroseconds"
                          label="RTTS in Microseconds"
                          :rules="[FORM_RULES.ruleRequiredField]"
                          type="Number"
                          :disabled="true"
                      />
                    </v-col>
                  </v-row>
                </v-form>
              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions class="pb-4">
            <v-row justify="center">
              <v-col
                v-if="!hasResults"
                cols="12"
                align="center"
              >
                <v-btn
                    @click="submitToExperiment"
                    color="primary"
                    :disabled="isLoading"
                >
                  Begin Experiment
                </v-btn>
              </v-col>
              <v-col
                  v-if="!hasResults"
                  cols="12"
                  align="center"
              >
                <input
                  ref="fileInput"
                  type="file"
                  accept="application/json"
                  @change="onFileChange"
                  hidden
                >
                <v-btn
                    @click="$refs.fileInput.click()"
                    color="brown lighten-2"
                    :disabled="isLoading"
                >
                  Load results from JSON file
                </v-btn>
              </v-col>
              <v-col
                v-if="hasResults"
                cols="12"
                align="center"
              >
                <v-btn
                    @click="submitToAnalysis"
                    color="primary"
                    :disabled="isLoading"
                >
                  Analyze
                </v-btn>
              </v-col>
              <v-col
                v-if="hasResults"
                cols="12"
                align="center"
              >
                <v-btn
                    @click="downloadResults"
                    color="light-green lighten-3"
                    :disabled="isLoading"
                >
                  Download Results
                </v-btn>
              </v-col>
            </v-row>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>

import download from 'downloadjs'

import Alert from "@/components/Alert"
import AnalyzerService from "@/services/analyzer_service";
import ExperimenterService from "@/services/experimenter_service"
import {FormRules} from "@/common/form_rules"
import LoadingBar from "@/components/LoadingBar"


export default {
  name: 'HomeView',
  components: {Alert, LoadingBar},
  data () {
    return {
      alertActive: false,
      alertMessage: '',
      alertMessageType: 'info',
      currentTime: Date.now(),
      isFormValid: false,
      isLoading: false,
      interactions: 100,
      batchSize: 10,
      label: '',
      intervalBetweenBatchesInMilliseconds: 0,
      startingTime: Date.now(),
      hasResults: false,
      mean: 0,
      standardDeviation: 0,
      failures: 0,
      rttsInMicroseconds: []
    }
  },
  computed: {
    totalTime () {
      let deltaTime = Math.abs(this.currentTime - this.startingTime) / 1000
      const hours = Math.floor(deltaTime / 3600)
      deltaTime -= hours * 3600
      const minutes = Math.floor(deltaTime / 60) % 60
      deltaTime -= minutes * 60
      const seconds = Math.floor(deltaTime % 60)
      return `${hours}h:${minutes}m:${seconds}s`
    },
    loadingMessage () {
      return `Uploading Scenes. Elapsed time: ${this.totalTime}`
    }
  },
  created () {
    this.EXPERIMENTER_SERVICE = new ExperimenterService()
    this.ANALYZER_SERVICE = new AnalyzerService()
    this.FORM_RULES = new FormRules()
  },
  methods: {
    clearAlert () {
      this.alertActive = false
      this.alertMessage = ''
      this.alertMessageType = 'info'
    },

    /**
     * @param {Object} data
     */
    updateResults (data) {
      this.hasResults = true
      this.label = data.label
      this.interactions = data.workload.interactions
      this.intervalBetweenBatchesInMilliseconds = data.workload.intervalBetweenBatchesInMilliseconds
      this.batchSize = data.workload.batchSize

      this.mean = data.mean
      this.standardDeviation = data.standardDeviation
      this.failures = data.failures
      this.rttsInMicroseconds = data.rttsInMicroseconds
    },

    buildDataObject () {
      return {
        label: this.label,
        workload: {
          interactions: this.interactions,
          intervalBetweenBatchesInMilliseconds: this.intervalBetweenBatchesInMilliseconds,
          batchSize: this.batchSize,
        },
        mean: this.mean,
        standardDeviation: this.standardDeviation,
        failures: this.failures,
        rttsInMicroseconds: this.rttsInMicroseconds
      }
    },

    downloadResults () {
      const jsonResults = this.buildDataObject()
      download(JSON.stringify(jsonResults), `${this.label}_results_interactions${this.interactions}_batchSize${this.batchSize}_interval${this.intervalBetweenBatchesInMilliseconds}.json`, "application/json")
    },

    /**
     * Begins the experiment.
     */
    submitToExperiment () {
      this.$refs.form.validate()
      this.clearAlert()

      if (this.isFormValid) {
        this.startingTime = Date.now()
        this.isLoading = true

        const parameters = {
          "interactions": Number(this.interactions),
          "batchSize": Number(this.batchSize),
          "intervalBetweenBatchesInMilliseconds": Number(this.intervalBetweenBatchesInMilliseconds)
        }

        const successCallBack = (response) => {
          this.alertMessage = 'Successfully executed the experiment'
          this.alertMessageType = 'success'
          response.data.label = this.label
          this.updateResults(response.data)
        }

        const errorCallBack = (error) => {
          this.alertMessage = 'Failed to run experiment.'
          this.alertMessageType = 'error'
        }

        const finallyCallBack = () => {
          this.isLoading = false
          this.alertActive = true
        }

        this.EXPERIMENTER_SERVICE.runExperiment(parameters, successCallBack, errorCallBack, finallyCallBack)
      }
    },

    /**
     * Begins the analysis.
     */
    submitToAnalysis () {
      this.$refs.form.validate()
      this.clearAlert()

      if (this.isFormValid) {
        this.startingTime = Date.now()
        this.isLoading = true

        const parameters = this.buildDataObject()

        const successCallBack = (response) => {
          this.alertMessage = 'Successfully executed the analysis'
          this.alertMessageType = 'success'
          console.log(response.data)
        }

        const errorCallBack = (error) => {
          this.alertMessage = 'Failed to run analysis.'
          this.alertMessageType = 'error'
        }

        const finallyCallBack = () => {
          this.isLoading = false
          this.alertActive = true
        }

        this.ANALYZER_SERVICE.analyze(parameters, successCallBack, errorCallBack, finallyCallBack)
      }
    },

    onFileChange (event) {
      let files = event.target.files || event.dataTransfer.files
      if (!files.length) return
      this.readFile(files[0])
    },

    readFile(file) {
      let reader = new FileReader()
      reader.onload = e => {
        let json = JSON.parse(e.target.result)
        this.updateResults(json)
      }
      reader.readAsText(file);
    }
  },
  watch: {
    isLoading (newIsLoading) {
      if (!newIsLoading) {
        console.log(`total time: ${this.totalTime}`)
      }
    },
    currentTime: {
      handler() {
        setTimeout(() => {
          this.currentTime = Date.now()
        }, 1000)
      },
      immediate: true
    }
  }
}
</script>

<style scoped>
  .alert-class {
    position: fixed;
    left: 50%;
    top: 50px;
    transform: translate(-50%, -50%);
    margin: 0 auto;
    width: 80%;
  }
</style>
