import React from "react";
import { DatePicker } from "antd";
import '../style/shortlink.css';

class Shortlink extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            getedorigin: "",
            getedorigin2: "",
            getedorigin3: "",
            getedshort: "",
            getedshort2: "",
            getedshort3: "",
            getedshort4: "",
            getedcomment: "",
            getedcomment2: "",
            getedcomment3: "",
            myform: "",
            linkid: "",
            shorted: "",
            isActive: "",
            deleteid: "",
            gotoinfo: 0,
            msg_create: "",
            formattedUTC1: "",
            formattedUTC2: "",
            updateTime1: "",
            updateTime2: ""
        };
    }

    handleDateChange1 = (date, dateString) => {
        this.setState({
            formattedUTC1: dateString
        });
    };

    handleDateChange2 = (date, dateString) => {
        this.setState({
            formattedUTC2: dateString
        });
    };

    getOrigin = (event) => {
        this.setState({
            getedorigin: event.target.value
        });
    };

    getShort = (event) => {
        this.setState({
            getedshort: event.target.value
        });
    };

    getComment = (event) => {
        this.setState({
            getedcomment: event.target.value
        });
    };

    getDeleteid = (event) => {
        this.setState({
            deleteid: event.target.value
        });
    };

    render() {
        if (0 === this.state.gotoinfo)
            return (
                <div className="mainbody">
                    <div className="text">
                        <div style={{ paddingLeft: 190, paddingBottom: 15 }}>
                            <h2 style={{ marginBottom: 5 }}>Shorten and personalize any link.</h2>
                        </div>
                        <div style={{ paddingLeft: 190 }}>
                            <h3>Get real-time traffic statistics for your links.</h3>
                        </div>
                        <div style={{ paddingLeft: 190, paddingBottom: 15 }}>
                            <h3>Free service.</h3>
                        </div>
                        <div className="contain">
                            <div className="case1">
                                <div className="caseinside">
                                    <p>Link to shorten</p>
                                    <input type="text" className="input-txt" id="linktoshort" onChange={this.getOrigin}></input>
                                    <p>Customize your short link:</p>
                                    http://localhost:8080/<input type="text" id="diy" className="input-txt" onChange={this.getShort}></input>
                                    <p>Optional comment</p>
                                    <input type="text" className="input-txt" id="commit" onChange={this.getComment}></input>
                                    <br />
                                    <p>Select start time:</p>
                                    <DatePicker onChange={this.handleDateChange1} />
                                    <p>Select end time:</p>
                                    <DatePicker onChange={this.handleDateChange2} />
                                    <br />
                                    <input
                                        type="button"
                                        className="input-btn"
                                        value="Shorten"
                                        onClick={() => {
                                            const { formattedUTC1, formattedUTC2 } = this.state;
                                            const formattedTime1 = `${formattedUTC1}T00:00:00Z`;
                                            const formattedTime2 = `${formattedUTC2}T23:59:59Z`;

                                            // 打印请求的详细信息
                                            console.log('Sending request to create short link:');
                                            console.log('URL: http://localhost:8080/api/link/create');
                                            console.log('Method: POST');
                                            console.log('Headers: {\'Content-Type\': \'application/json\'}');
                                            console.log('Body:', JSON.stringify({
                                                "short": this.state.getedshort,
                                                "origin": this.state.getedorigin,
                                                "comment": this.state.getedcomment,
                                                "start_time": formattedTime1,
                                                "end_time": formattedTime2
                                            }));

                                            fetch("http://localhost:8080/api/link/create", {
                                                credentials: 'include',
                                                method: "POST",
                                                headers: {
                                                    'Content-Type': 'application/json',
                                                },
                                                body: JSON.stringify({
                                                    "short": this.state.getedshort,
                                                    "origin": this.state.getedorigin,
                                                    "comment": this.state.getedcomment,
                                                    "start_time": formattedTime1,
                                                    "end_time": formattedTime2
                                                })
                                            })
                                            .then(response => {
                                                console.log('Create response:', response);
                                                return response.json();
                                            })
                                            .then(data => {
                                                console.log('Create response data:', data);
                                                if (data.data) {
                                                    this.setState({
                                                        msg_create: data.msg,
                                                        shorted: data.data.short
                                                    });
                                                } else {
                                                    this.setState({
                                                        msg_create: data.msg
                                                    });
                                                }
                                                document.getElementById('generated').value = data.data ? data.data.short : "";
                                            })
                                            .catch(error => console.error('Error creating link:', error));
                                        }}
                                    />
                                </div>
                            </div>

                            <div className="case1">
                                <div>
                                    <p>Generated Link</p><br />
                                    <p>Your Short Link:</p>
                                    <input type="text" className="input-txt" id="generated" readOnly></input>
                                    <br />
                                    <p id="shortenmsg" style={{ fontSize: 20, color: 'red' }}>{this.state.msg_create}</p>
                                </div>
                            </div>

                            <div className="case">
                                <div>
                                    Input the short link you want to DELETE:<br />
                                    <input type="text" className="input-txt" id="delete" onChange={this.getDeleteid}></input><br />
                                    <input
                                        type="button"
                                        className="input-btn"
                                        value="Delete"
                                        onClick={() => {
                                            console.log('Sending request to delete short link:');
                                            console.log('URL: http://localhost:8080/api/link/delete');
                                            console.log('Method: POST');
                                            console.log('Headers: {\'Content-Type\': \'application/json\'}');
                                            console.log('Body:', JSON.stringify({
                                                "short": this.state.deleteid
                                            }));

                                            fetch("http://localhost:8080/api/link/delete", {
                                                method: "POST",
                                                credentials: 'include',
                                                headers: {
                                                    'Content-Type': 'application/json',
                                                },
                                                body: JSON.stringify({
                                                    "short": this.state.deleteid
                                                })
                                            })
                                            .then(response => {
                                                console.log('Delete response:', response);
                                                return response.json();
                                            })
                                            .then(data => {
                                                console.log('Delete response data:', data);
                                                document.getElementById('deletemsg').innerText = data.msg;
                                            })
                                            .catch(error => console.error('Error deleting link:', error));
                                        }}
                                    />
                                    <p id="deletemsg" style={{ fontSize: 20, color: 'red' }}></p>
                                </div>
                            </div>

                            <div className="case1">
                                <div>
                                    <p>Modify Link Comment:</p>
                                    <input type="text" className="input-txt" id="comment" onChange={this.getComment3}></input>
                                    <input type="text" className="input-txt" id="short" onChange={this.getShort4}></input>
                                    <input
                                        type="button"
                                        className="input-btn"
                                        value="Update"
                                        onClick={() => {
                                            console.log('Sending request to update short link:');
                                            console.log('URL: http://localhost:8080/api/link/update');
                                            console.log('Method: POST');
                                            console.log('Headers: {\'Content-Type\': \'application/json\'}');
                                            console.log('Body:', JSON.stringify({
                                                "short": this.state.getShort4,
                                                "comment": this.state.getComment3
                                            }));

                                            fetch("http://localhost:8080/api/link/update", {
                                                method: "POST",
                                                credentials: 'include',
                                                headers: {
                                                    'Content-Type': 'application/json',
                                                },
                                                body: JSON.stringify({
                                                    "short": this.state.getShort4,
                                                    "comment": this.state.getComment3
                                                })
                                            })
                                            .then(response => {
                                                console.log('Update response:', response);
                                                return response.json();
                                            })
                                            .then(data => {
                                                console.log('Update response data:', data);
                                                document.getElementById('updatemsg').innerText = data.msg;
                                            })
                                            .catch(error => console.error('Error updating link:', error));
                                        }}
                                    />
                                    <p id="updatemsg" style={{ fontSize: 20, color: 'red' }}></p>
                                </div>
                            </div>

                            <div className="case">
                                <div>
                                    <p>Check Short Link Info:</p>
                                    <input type="text" className="input-txt" id="check" onChange={this.getShort2}></input><br />
                                    <input
                                        type="button"
                                        className="input-btn"
                                        value="Check"
                                        onClick={() => {
                                            console.log('Sending request to check short link info:');
                                            console.log('URL: http://localhost:8080/api/link/info');
                                            console.log('Method: POST');
                                            console.log('Headers: {\'Content-Type\': \'application/json\'}');
                                            console.log('Body:', JSON.stringify({
                                                "short": this.state.getShort2
                                            }));

                                            fetch("http://localhost:8080/api/link/info", {
                                                method: "POST",
                                                credentials: 'include',
                                                headers: {
                                                    'Content-Type': 'application/json',
                                                },
                                                body: JSON.stringify({
                                                    "short": this.state.getShort2
                                                })
                                            })
                                            .then(response => {
                                                console.log('Info response:', response);
                                                return response.json();
                                            })
                                            .then(data => {
                                                console.log('Info response data:', data);
                                                if (data.data) {
                                                    document.getElementById('shortinfo').innerText = data.data.short;
                                                    document.getElementById('origininfo').innerText = data.data.origin;
                                                    document.getElementById('commentinfo').innerText = data.data.comment;
                                                    document.getElementById('starttime').innerText = data.data.start_time;
                                                    document.getElementById('endtime').innerText = data.data.end_time;
                                                } else {
                                                    this.setState({
                                                        msg_create: data.msg
                                                    });
                                                }
                                            })
                                            .catch(error => console.error('Error fetching link info:', error));
                                        }}
                                    />
                                    <p id="shortinfo" style={{ fontSize: 20, color: 'red' }}></p>
                                    <p id="origininfo" style={{ fontSize: 20, color: 'blue' }}></p>
                                    <p id="commentinfo" style={{ fontSize: 20, color: 'blue' }}></p>
                                    <p id="starttime" style={{ fontSize: 20, color: 'blue' }}></p>
                                    <p id="endtime" style={{ fontSize: 20, color: 'blue' }}></p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            );
    }
}

export default Shortlink;
