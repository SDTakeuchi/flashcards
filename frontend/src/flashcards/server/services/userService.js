export const userService = {
  authenticate,
};

function authenticate(username, password) {
  console.log("let's go")
  console.log("username: ", username);
  console.log("password: ", password);
  if (username !== "admin" && password !== "admin") {
    return null;
  }

  const user = {
    id: "9001",
    name: "Web Admin",
    email: "admin@example.com",
  };

  return user;
}
