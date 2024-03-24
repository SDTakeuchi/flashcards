import { FlashCard } from "@/component/card";
import { getServerAuthSession } from "@/server/auth";
import Header from "@/component/header";

export default async function Cards() {
  const authSession = await getServerAuthSession();

  console.log("authSession: ", authSession);

  return (
    <div>
      <Header />
      <FlashCard />
    </div>
  );
}
