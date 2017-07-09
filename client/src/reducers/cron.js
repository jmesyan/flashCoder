import React from 'react'


const initialState = {
    messages: [{type: 'MESSAGE',
                text: 'Welcome to our chatting room!'}],
};

export default function cron(state = initialState, action) {
    switch (action.type) {
        case 'WELCOME':
            const message = action.message;
            return { ...state,
                messages: [ ...state , {
                    userName: 'good'
                }]
            }
        default:
            return state;
    }
}
