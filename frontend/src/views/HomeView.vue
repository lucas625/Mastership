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
                  Run the experiment
                </h1>
              </v-col>
            </v-row>
          </v-card-title>
          <v-card-text>
            <v-row justify="center">
              <v-col align="center">
                <v-form v-model="isFormValid">
                  <v-row>
                    <v-col>
                      <v-text-field
                          v-model="interactions"
                          label="Number of requests"
                          :rules="[FORM_RULES.ruleRequiredField, FORM_RULES.ruleMinMaxValueField(100, 100000)]"
                          type="Number"
                      />
                      <v-text-field
                          v-model="batchSize"
                          label="Size of groups of requests"
                          :rules="[FORM_RULES.ruleRequiredField, FORM_RULES.ruleMinMaxValueField(1, 10000)]"
                          type="Number"
                      />
                      <v-text-field
                          v-model="intervalBetweenBatchesInMilliseconds"
                          label="Interval in ms between groups of requests"
                          :rules="[FORM_RULES.ruleMinMaxValueField(0, 1000)]"
                          type="Number"
                      />
                    </v-col>
                  </v-row>
                </v-form>
              </v-col>
            </v-row>
          </v-card-text>
          <v-card-actions>
            <v-row justify="center">
              <v-col align="center">
                <v-btn
                    @click="submit"
                    color="primary"
                    :disabled="isLoading || !isFormValid"
                >
                  Begin Experiment
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

import RayTracingControllerService from "@/services/controller_service"
import {FormRules} from "@/common/form_rules";
import LoadingBar from "@/components/LoadingBar";
import Alert from "@/components/Alert";


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
      intervalBetweenBatchesInMilliseconds: 0,
      startingTime: Date.now()
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
    this.CONTROLLER_SERVICE = new RayTracingControllerService()
    this.FORM_RULES = new FormRules()
  },
  methods: {
    clearAlert () {
      this.alertActive = false
      this.alertMessage = ''
      this.alertMessageType = 'info'
    },

    /**
     * Begins the experiment.
     */
    submit () {
      this.startingTime = Date.now()
      this.isLoading = true
      this.clearAlert()

      const parameters = {
        "interactions": this.interactions,
        "batchSize": this.batchSize,
        "intervalBetweenBatchesInMilliseconds": this.intervalBetweenBatchesInMilliseconds
      }

      const successCallBack = (response) => {
        this.alertMessage = 'Successfully executed the experiment'
        this.alertMessageType = 'success'
      }

      const errorCallBack = (error) => {
        this.alertMessage = 'Failed to run experiment.'
        this.alertMessageType = 'error'
      }

      const finallyCallBack = () => {
        this.isLoading = false
        this.alertActive = true
      }

      this.CONTROLLER_SERVICE.runExperiment(parameters, successCallBack, errorCallBack, finallyCallBack)
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
