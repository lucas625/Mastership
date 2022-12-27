export class URLHelper {
    /**
     * Builds the base url. If the selected url is not empty, then returns it, else it gets the current origin.
     * @param {string} url
     * @returns {string}
     */
    static buildBaseURL(url) {
        return url ? url : window.location.origin
    }
}