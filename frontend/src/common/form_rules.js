export class FormRules {
    /**
     * Forces the field to be present.
     * @param value
     * @returns {boolean|string}
     */
    ruleRequiredField (value) {
        const msg = 'This field is required'
        let is_valid
        if (typeof value === 'number') {
            is_valid = Boolean(value)
        } else {
            is_valid = value && value.length > 0
        }
        return is_valid || msg
    }

    /**
     * Forces the form field to provide a value between the min and the max.
     * @param {Number} min
     * @param {Number} max
     * @return {function}
     */
    ruleMinMaxValueField (min, max) {
        return function(value) {
            let msg
            let is_valid = true
            if (value < min) {
                is_valid = false
                msg = `The value must be at least ${min}.`
            } else if (value > max) {
                is_valid = false
                msg = `The value must not be greater than ${max}.`
            }
            return is_valid || msg
        }
    }
}