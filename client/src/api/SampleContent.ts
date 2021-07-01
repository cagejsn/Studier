import Axios from "axios";
import { SampleContent  } from "@/models/LandingPage";

const axios = Axios.create({ baseURL: "https://localhost:8091" });

export const getSampleContent = (sampleContentType: string) => {
  return axios.get<SampleContent>('/sample-content', { params: { sampleContentType, userID: 1} });
};
