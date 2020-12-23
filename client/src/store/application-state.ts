import { Problem } from '@/models/Problem';


interface ApplicationState {
    problemBeingViewed?: Problem
}



export const initialState: ApplicationState = {
    problemBeingViewed: undefined
}