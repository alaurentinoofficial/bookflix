import axios from "axios";

const baseDomain = "http://api.bookflix.laurentino.me";
const baseURL = `${baseDomain}`;


export default axios.create({
  baseURL,
  headers: {
  }
});