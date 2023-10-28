import { useState } from "react";
import { ISignUpInfo, useSignUpMutation } from "../../core/api/authApi";
import "./SignUp.scss";
import { ChangeEvent } from "../../App.typig";

function SignUpForm() {
  const [signUp] = useSignUpMutation();

  const [formData, setFormData] = useState<ISignUpInfo>({
    userName: "",
    password: "",
    email: "",
    fullName: "",
  });
  const [message, setMessage] = useState("");

  function handleChange(event: ChangeEvent) {
    const { id, value } = event.target;
    setFormData((prevState) => ({ ...prevState, [id]: value }));
  }

  function clickSignUp() {
    if (
      formData.password &&
      formData.email &&
      formData.userName &&
      formData.fullName
    ) {
      signUp(formData);
    } else [setMessage("Введите все данные")];
  }

  return (
    <div className="signup">
      <h3>Login</h3>

      <label htmlFor="fullName">FullName</label>
      <input
        type="text"
        placeholder="vladimir kabanets"
        id="fullName"
        value={formData.fullName}
        onChange={handleChange}
      />

      <label htmlFor="email">Email</label>
      <input
        type="text"
        placeholder="capvok@yandex.ru"
        id="email"
        value={formData.email}
        onChange={handleChange}
      />

      <label htmlFor="userName">Username</label>
      <input
        type="text"
        placeholder="capvok"
        id="userName"
        value={formData.userName}
        onChange={handleChange}
      />

      <label htmlFor="password">Password</label>
      <input
        type="password"
        placeholder="123123"
        id="password"
        value={formData.password}
        onChange={handleChange}
      />

      <button onClick={clickSignUp}>Log In</button>
      <p>{message}</p>
    </div>
  );
}

export { SignUpForm };
