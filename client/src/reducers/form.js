import { createForms } from 'react-redux-form';

const initCronForm = {
    Second:'',
    Minute:'',
    Hour:'',
    Day:'',
    Month:'',
    Week:''
}

const initBehavior = {
    bname:'',
    operate:'',
    paramsList:'',
    remark:''
}

const initOperate = {
    opname:'',
    optag:'',
    remark:''
}




var reduceForm = createForms({
    cronForm: initCronForm,
    behaviorForm:initBehavior,
    operateForm:initOperate
})

export default reduceForm; 