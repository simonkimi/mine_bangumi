import axios from "axios";

const instance = axios.create({
  timeout: 10000,
  baseURL: "http://localhost:7962",
  headers: { "Content-Type": "application/json" },
});

export default instance;
