import "./AuthPage.scss";
import { useState } from "react";

import { SignUpForm, LogInForm } from "../../componets";

export const AuthPage: React.FC = () => {
  const [isLogIn, setIsLogIn] = useState(true);

  return (
    <div className="auth_page">
      <div className="form_block">
        {isLogIn ? <LogInForm /> : <SignUpForm />}
      </div>
      <div className="image_block">
        <div className="image"></div>
        <button
          className="btn-change_is_login"
          onClick={() => setIsLogIn(!isLogIn)}
        >
          {isLogIn ? "У меня нет аккаунта" : "Я меня есть аккаунт"}
        </button>
      </div>
    </div>
  );
};
