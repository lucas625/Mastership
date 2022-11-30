import axios from 'axios'

/**
 * Access layer to the controller service.
 * @constructor
 */
export default class ControllerService {
  /**
   * {ControllerService} constructor.
   */
  constructor () {
    this.client = axios.create({ baseURL: `${process.env.VUE_APP_MSC_EXPERIMENTER_URL}` })
  }

  /**
   * Runs the experiment.
   * @param {object} data - The parameters to the analysis.
   * @param {function} successCallBack - The function to be performed on success.
   * @param {function} errorCallback - The function to be performed on error.
   * @param {function} finallyCallback - The function to be performed after the success/error callback.
   */
  runExperiment (data, successCallBack, errorCallback, finallyCallback) {
    // TODO: Find way to enable the request on backend side
    this.client.post('experiment', data, null)
      .then(successCallBack)
      .catch(errorCallback)
      .then(finallyCallback)
  }
}
