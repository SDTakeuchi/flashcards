import axios from "axios";
// import { cookies } from "next/headers";
import Cookies from "js-cookie";

export const userService = {
  authenticate,
};

function authenticate(username, password) {
  let user = {
    id: undefined,
    name: undefined,
    email: null,
    accessToken: undefined,
  };

  axios
    .post("http://flashcard_backend:8000/backend_api/auth/login", {
      name: username,
      password: password,
    })
    .then((response) => {
      console.log("response: ", response);
      user.id = response.data?.user_id;
      user.accessToken = response.data?.access_token;
      Cookies.set("access_token", response.data.access_token);
    })
    .catch((error) => {
      console.log(error);
    });

  return user;
}
