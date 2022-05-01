import logo from "./logo.svg"
import "./App.css"
import { initializeApp } from "firebase/app"
import {
  getAuth,
  signInWithCustomToken,
  signInWithEmailAndPassword,
  GoogleAuthProvider,
  signInWithPopup,
} from "firebase/auth"
import { useEffect, useState } from "react"
import axios from "axios"

function App() {
  var firebaseConfig = {
    apiKey: "AIzaSyB1KmtYSlkCoklqcUcgcOVu2wAzgbj71Mw",
    authDomain: "is766-project.firebaseapp.com",
    projectId: "is766-project",
    storageBucket: "is766-project.appspot.com",
    messagingSenderId: "36177987727",
    appId: "1:36177987727:web:aba1e56b311a6cf38448fa",
    measurementId: "G-XS5M2305JW",
  }
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [user, setUser] = useState("")
  const [idToken, setIdToken] = useState("")
  const app = initializeApp(firebaseConfig)
  const provider = new GoogleAuthProvider()

  const auth = getAuth(app)
  const loginFunction = async () => {
    const login = await signInWithEmailAndPassword(auth, username, password)
    setUser(login)
    console.log(login)
  }
  useEffect(() => {}, [])

  const handleLoginWithGoogle = async () => {
    const result = await signInWithPopup(auth, provider)

    // This gives you a Google Access Token. You can use it to access the Google API.
    const credential = GoogleAuthProvider.credentialFromResult(result)
    const token = credential.accessToken
    // The signed-in user info.
    const user = result.user
    setUser(user)
    console.log(result.user)

    const res = await axios.post("http://localhost:8080/register", {
      email: result.user.email,
      firstname: result.user.displayName.split(" ")[0],
      lastname:
        result.user.displayName.split(" ").length >= 2
          ? result.user.displayName.split(" ")[1]
          : null,
      phone: user.phoneNumber,
    })
    console.log(res.data)
  }

  return (
    <div className="App">
      <meta
        name="google-signin-client_id"
        content="YOUR_CLIENT_ID.apps.googleusercontent.com"
      ></meta>
      <header className="App-header">
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            // alignItems: "center",
            width: "256px",
          }}
        >
          <div
            style={{
              display: "flex",
              margin: "5px",
              justifyContent: "end",
            }}
          >
            <label
              style={{ marginRight: "3px", textAlign: "end" }}
              htmlFor="username"
            >
              email:{" "}
            </label>
            <input
              type="text"
              name="username"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div
            style={{
              display: "flex",
              margin: "5px",
              justifyContent: "end",
            }}
          >
            <label
              style={{ marginRight: "3px", textAlign: "end" }}
              htmlFor="password"
            >
              password:{" "}
            </label>
            <input
              type="password"
              name="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <button style={{ width: "16rem" }} onClick={() => loginFunction()}>
            Login
          </button>
          <br />
          <div id="gSignInWrapper">
            <span class="label">Sign in with:</span>
            <div
              id="customBtn"
              class="customGPlusSignIn"
              onClick={() => handleLoginWithGoogle()}
            >
              <span class="icon"></span>
              <span class="buttonText">Google</span>
            </div>
          </div>
        </div>
        <div>
          {user?.user?.email ? (
            <p style={{ width: "50vw", wordWrap: "break-word" }}>
              {user?.user?.email}
            </p>
          ) : (
            <></>
          )}
          {user?._tokenResponse?.idToken ? (
            <p style={{ width: "50vw", wordWrap: "break-word" }}>
              {user?._tokenResponse?.idToken}
            </p>
          ) : (
            <></>
          )}
          {user.displayName ? (
            <p style={{ width: "50vw", wordWrap: "break-word" }}>
              {user?.displayName}
            </p>
          ) : (
            <></>
          )}
          {user.accessToken ? (
            <p style={{ width: "50vw", wordWrap: "break-word" }}>
              {user?.accessToken}
            </p>
          ) : (
            <></>
          )}
        </div>
      </header>
    </div>
  )
}

export default App
