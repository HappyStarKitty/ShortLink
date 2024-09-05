import React from "react";
import Register from "./register";
import Shortlink from "./shortlink";
import "../style/loginandregister.css";

class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            getedname: "",
            getedpwd: "",
            getedmail: "",
            responcode: 0,
            responjwt: "",
            catchaUrl: "",
            catchaid: "",
            getedCap: "",
            respon: 0,
            message: ""
        };
    }

    componentDidMount() {
        // Fetch initial captcha URL after component mounts
        this.fetchCatchaUrl();
    }

    fetchCatchaUrl() {
        console.log("Fetching captcha URL...");
        fetch("http://localhost:8080/api/user/captcha", {
            method: "GET",
            credentials: 'include',
        })
        .then(response => {
            console.log("Captcha URL fetch response:", response);
            return response.json();
        })
        .then(data => {
            console.log("Captcha URL fetch data:", data);
            this.setState({
                catchaUrl: data.captcha_url,
                catchaid: data.captcha_id
            });
        })
        .catch(error => {
            console.error("Captcha URL fetch error:", error);
        });
    }

    getName = (event) => {
        this.setState({
            getedname: event.target.value
        });
    }

    getMail = (event) => {
        this.setState({
            getedmail: event.target.value
        });
    }

    getPwd = (event) => {
        this.setState({
            getedpwd: event.target.value
        });
    }

    getCaptcha = (event) => {
        this.setState({
            getedCap: event.target.value
        });
    }

    handleClick() {
        console.log("Logging out...");
        fetch("http://localhost:8080/api/user/logout", {
            method: "POST",
            credentials: 'include',
        })
        .then(response => {
            console.log("Logout response:", response);
            return response.json();
        })
        .then(data => {
            console.log("Logout data:", data);
            this.setState({
                responcode: 0
            });
        })
        .catch(error => {
            console.error("Logout error:", error);
        });
    }

    render() {
        if (0 === this.state.responcode)
            return (
                <div className="container">
                    <div className="box">
                        <div className="title">Register</div>
                        <Register />
                    </div>

                    <div className="box">
                        <img
                            src={this.state.catchaUrl}
                            alt="Captcha"
                            className="catcha"
                            onClick={() => this.fetchCatchaUrl()}
                        />
                        <div className="title">Login</div>
                        <input
                            type="email"
                            name="email"
                            id="logmail"
                            className="input-text"
                            placeholder="email"
                            autoComplete="email"
                            onChange={this.getMail}
                        /><br />
                        <input
                            type="password"
                            name="password"
                            id="logpwd"
                            className="input-text"
                            placeholder="password"
                            autoComplete="current-password"
                            onChange={this.getPwd}
                        /><br />
                        <input
                            type="text"
                            name="captcha"
                            id="logcap"
                            className="input-text"
                            placeholder="captcha"
                            autoComplete="off"
                            onChange={this.getCaptcha}
                        /><br />
                        <input
                            type="button"
                            className="input-button"
                            value="LOGIN"
                            onClick={() => {
                                this.fetchCatchaUrl();
                                console.log("Logging in...");
                                fetch("http://localhost:8080/api/user/login", {
                                    method: "POST",
                                    credentials: 'include',
                                    headers: { 'Content-Type': 'application/json' },
                                    body: JSON.stringify({
                                        "email": this.state.getedmail,
                                        "password": this.state.getedpwd,
                                        "captcha_id": this.state.catchaid,
                                        "captcha_value": this.state.getedCap
                                    })
                                })
                                .then(response => {
                                    console.log("Login response:", response);
                                    return response.json();
                                })
                                .then(data => {
                                    console.log("Login data:", data);
                                    this.setState({
                                        responcode: data.code,
                                        message: data.msg,
                                        respon: 1,
                                    });
                                })
                                .catch(error => {
                                    console.error("Login error:", error);
                                });
                            }}
                        />
                    </div>
                </div>
            );
        else if (0 !== this.state.responcode) {
            return (
                <Shortlink onClick={() => this.handleClick()} responjwt={this.state.responjwt} />
            );
        }
        else {
            return (
                <div className="container">
                    <div className="box">
                        <div className="title">Register</div>
                        <Register />
                    </div>
                    <div className="box">
                        <img
                            src={this.state.catchaUrl}
                            alt="Captcha"
                            className="catcha"
                            onClick={() => this.fetchCatchaUrl()}
                        />
                        <div className="title">Login</div>
                        <input
                            type="email"
                            name="email"
                            id="logmail"
                            className="input-text"
                            placeholder="email"
                            autoComplete="email"
                            onChange={this.getMail}
                        /><br />
                        <input
                            type="password"
                            name="password"
                            id="logpwd"
                            className="input-text"
                            placeholder="password"
                            autoComplete="current-password"
                            onChange={this.getPwd}
                        /><br />
                        <input
                            type="text"
                            name="captcha"
                            id="logcap"
                            className="input-text"
                            placeholder="captcha"
                            autoComplete="off"
                            onChange={this.getCaptcha}
                        /><br />
                        <input
                            type="button"
                            className="input-button"
                            value="LOGIN"
                            onClick={() => {
                                this.fetchCatchaUrl();
                                console.log("Logging in...");
                                fetch("http://localhost:8080/api/user/login", {
                                    method: "POST",
                                    credentials: 'include',
                                    headers: { 'Content-Type': 'application/json' },
                                    body: JSON.stringify({
                                        "email": this.state.getedmail,
                                        "password": this.state.getedpwd,
                                        "captcha_id": this.state.catchaid,
                                        "captcha_value": this.state.getedCap
                                    })
                                })
                                .then(response => {
                                    console.log("Login response:", response);
                                    return response.json();
                                })
                                .then(data => {
                                    console.log("Login data:", data);
                                    this.setState({
                                        responcode: data.code,
                                        message: data.msg,
                                        respon: 1,
                                    });
                                })
                                .catch(error => {
                                    console.error("Login error:", error);
                                });
                            }}
                        />
                        <p style={{ display: this.state.respon ? 'flex' : 'none', color: 'red' }}>{this.state.message}</p>
                    </div>
                </div>
            );
        }
    }
}

export default Login;
