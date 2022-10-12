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
    // TODO: Find way to add path on development as the way it is, it will only work for docker/kubernetes
    this.client = axios.create({ baseURL: `api/experiment-controller/` })
  }

  /**
   * Runs the experiment.
   * @param {object} data - The parameters to the analysis.
   * @param {function} successCallBack - The function to be performed on success.
   * @param {function} errorCallback - The function to be performed on error.
   * @param {function} finallyCallback - The function to be performed after the success/error callback.
   */
  runExperiment (data, successCallBack, errorCallback, finallyCallback) {
    this.client.post('run-experiment', data)
      .then(successCallBack)
      .catch(errorCallback)
      .then(finallyCallback)
  }
}
