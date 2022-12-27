import axios from 'axios'
import {URLHelper} from "@/common/url_helper";

/**
 * Access layer to the analyzer service.
 * @constructor
 */
export default class AnalyzerService {
  /**
   * {AnalyzerService} constructor.
   */
  constructor () {
    this.client = axios.create({ baseURL: URLHelper.buildBaseURL(process.env.VUE_APP_MSC_ANALYZER_URL) })
  }

  /**
   * Runs the experiment.
   * @param {object} data - The parameters to the analysis.
   * @param {function} successCallBack - The function to be performed on success.
   * @param {function} errorCallback - The function to be performed on error.
   * @param {function} finallyCallback - The function to be performed after the success/error callback.
   */
  analyze (data, successCallBack, errorCallback, finallyCallback) {
    this.client.post('api/analyzer/analyze', data, null)
      .then(successCallBack)
      .catch(errorCallback)
      .then(finallyCallback)
  }
}
