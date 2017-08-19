import { createForms } from 'react-redux-form';

const initCronForm = {
    Second:'*',
    Minute:'*',
    Hour:'*',
    Day:'*',
    Month:'*',
    Week:'*'
}



var reduceForm = createForms({
    cronForm: initCronForm,
})

export default reduceForm; 