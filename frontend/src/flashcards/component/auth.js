import { useRouter } from "next/router";
import Cookies from "js-cookie";

const Auth = ({ children }) => {
  //router
  const router = useRouter();

  //Cookieのチェック（これをいろいろ認証タイプにより変更）
  const isSignedIn = Cookies.get("flashcard_access_token") !== undefined; // TODO: this is not secure

  if (!isSignedIn) {
    router.replace("/login");
  }

  return children;
};

export default Auth;
