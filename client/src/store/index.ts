import Vue from 'vue'
import Vuex from 'vuex'

import { initialState } from "./application-state"

import {updateProblem} from "@/api/UpdateProblem"

Vue.use(Vuex)

const store = new Vuex.Store({
    state: initialState,
    actions: {
        updateAnswers({commit}, answers){
            commit('setProblemAnswers', answers)

        },
        saveProblem({commit}){
            updateProblem(this.state.openProblem!)
            
        },
        updateSolution({commit}, solution){
            commit('setProblemSolution', solution)

        },
        updateProblemState({commit}, newState){
            commit('setProblemState', newState)
        }

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
        },
        setProblemAnswers(state, value){
            state.openProblem!.answers = value
        },
        setProblemSolution(state, value){
            state.openProblem!.solution = value
        },
        setProblemState(state, value){
            state.openProblem!.state = value
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
        },
        getAnswers(state){
            return state.openProblem?.answers
        },
        getSolution(state){
            return state.openProblem?.solution
        },
        getProblemState(state){
            return state.openProblem?.state
        }

    }
})

export default store;