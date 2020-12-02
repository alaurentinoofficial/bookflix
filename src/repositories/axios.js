import axios from "axios";

const baseDomain = "https://api.bookflix.laurentino.me";
const baseURL = `${baseDomain}`;


export default axios.create({
  baseURL,
  headers: {
  }
});