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
            console.log("saving problem", this.state.openProblem)
            updateProblem(this.state.openProblem!)
            
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
            console.log("setting answers to ", value)
            state.openProblem!.answers = value
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
        }

    }
})

export default store;