import Vue from 'vue'
import Vuex from 'vuex'

import { initialState } from "./application-state"

import {updateProblem} from "@/api/UpdateProblem"
import {createProblemAttempt, getProblemAttemptById, updateProblemAttempt} from '@/api/ProblemAttempt/ProblemAttempt'
import { ProblemAttemptSubmissionState } from '@/models/ProblemAttempt'
import { ProblemSolution, SolutionType } from '@/models/Problem'

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
        updateSolution({commit, state}, solution){
            commit('setProblemSolution', solution)
        },
        updateSolutionExplanation({commit}, solutionExplanation){
            commit('setProblemSolutionExplanation', solutionExplanation)
        },
        updateProblemState({commit}, newState){
            commit('setProblemState', newState)
        },

        //problem-attempt
        beginAttemptProblem({commit}, problemId){

            createProblemAttempt(problemId)
            .then(({data}) => commit('setProblemAttempt', data))
            .catch(error => console.log(error))

        },
        updateProblemAttempt({commit}, userInput) {

            commit('setProblemAttemptUserInput', userInput)

            updateProblemAttempt(this.state.currentAttempt!)
            .then(({data}) => commit('setProblemAttempt', data))
            .catch(error => console.log(error))

        },
        submitProblemAttempt({ commit }) {
            updateProblemAttempt(
                {
                    ...this.state.currentAttempt!,
                    submissionState: ProblemAttemptSubmissionState.SUBMITTED_PENDING_GRADE,

                }
            ).then(({ data }) => commit('setProblemAttempt', data))
                .catch(error => console.log(error))
        }



    },
    mutations: {
        setOpenProblem(state, problem) {
            if(!problem.type){
                problem.type = {
                    statementType: 0,
                    answersType: 0,
                    solutionType: 0,
                }
            }
            if(!problem.solution){
                problem.solution = {
                    solution: null,
                    explanation: null
                }
            }
            state.openProblem = problem
        },
        
        clearOpenProblem(state){
            state.openProblem = undefined
        },


        //type
        setProblemStatementType(state, value){
            state.openProblem!.type.statementType = value;
        },
        setProblemAnswersType(state, value){
            state.openProblem!.type.answersType = value;
        },
        setProblemSolutionType(state, value){
            state.openProblem!.type.solutionType = value;
        },


        setProblemAnswers(state, value){
            state.openProblem!.answers = value
        },
        setProblemSolution(state, value){
            state.openProblem!.solution.solution = value
        },
        setProblemSolutionExplanation(state, value){

            state.openProblem!.solution.explanation = value
        },

        setProblemState(state, value){
            state.openProblem!.state = value
        },

        //problem-attempt
        setProblemAttempt(state, value){
            state.currentAttempt = value
        },
        setProblemAttemptSubmissionState(state, value){
            state.currentAttempt!.submissionState = value
        },
        setProblemAttemptUserInput(state, value){
            state.currentAttempt!.userInput = value
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
        },
        getProblemVersion(state){
            return state.openProblem?.version
        },

        //problem-attempt
        getProblemAttempt(state){
            return state.currentAttempt
        },
        getProblemAttemptUserInput(state){
            return state.currentAttempt?.userInput
        },
        getProblemAttemptSubmissionState(state){
            return state.currentAttempt?.submissionState
        }

    }
})

export default store;