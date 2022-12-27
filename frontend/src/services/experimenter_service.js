import axios from 'axios'
import {URLHelper} from "@/common/url_helper";

/**
 * Access layer to the experimenter service.
 * @constructor
 */
export default class ExperimenterService {
  /**
   * {ExperimenterService} constructor.
   */
  constructor () {
    this.client = axios.create({ baseURL: URLHelper.buildBaseURL(process.env.VUE_APP_MSC_EXPERIMENTER_URL) })
  }

  /**
   * Runs the experiment.
   * @param {object} data - The parameters to the analysis.
   * @param {function} successCallBack - The function to be performed on success.
   * @param {function} errorCallback - The function to be performed on error.
   * @param {function} finallyCallback - The function to be performed after the success/error callback.
   */
  runExperiment (data, successCallBack, errorCallback, finallyCallback) {
    this.client.post('api/experimenter/experiment', data, null)
      .then(successCallBack)
      .catch(errorCallback)
      .then(finallyCallback)
  }
}
