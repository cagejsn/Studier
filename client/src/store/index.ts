import Vue from 'vue'
import Vuex from 'vuex'

import { initialState } from "./application-state"

Vue.use(Vuex)

const store = new Vuex.Store({
    state: initialState,
    actions: {
        

    },
    mutations: {
        setOpenProblem(state, problem) {
            if(!problem.type){
                problem.type = {
                    statementType: 0,
                    answersType: 0,
                }
            }
            state.openProblem = problem
        },
        
        clearOpenProblem(state){
            state.openProblem = undefined
        },

        setProblemStatementType(state, value){
            state.openProblem!.type.statementType = value;
        },
        setProblemAnswersType(state, value){
            state.openProblem!.type.answersType = value;
        }

    },
    getters: {

        getOpenProblem(state){
            // console.log("getter for problem")
            // if(state.openProblem){
                return state.openProblem
            // } else {
                // return {};
            // };
        }

    }
})

export default store;