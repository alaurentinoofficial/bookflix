import axios from "axios";

const baseDomain = "http://ec2-54-207-189-235.sa-east-1.compute.amazonaws.com:8080";
const baseURL = `${baseDomain}`;


export default axios.create({
  baseURL,
  headers: {
  }
});