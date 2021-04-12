import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react";

let endpoint = {window:process.env["REACT_ENDPOINT"]};

console.log(endpoint);

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      task: "",
      items: [],
      done: ""
    };
  }

  componentDidMount() {
    this.getTask();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { task } = this.state;
    if (task) {
      axios
        .post(
          endpoint + "/api/task",
          {
            task
          },
          {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        )
        .then(res => {
          this.getTask();
          this.setState({
            task: ""
          });
        });
    }
  };

  getTask = () => {
    axios.get(endpoint + "/api/task").then(res => {
      const data = res.data.response;
      if (data) {
        this.setState({
          items: data.map(item => {
            let color = "yellow";

            if (item.status) {
              color = "green";
            } 
            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word", margin: "10px 0px 0px 20px"}}>{item.task}</div>
                  </Card.Header>

                  <Card.Content extra textAlign="right">
                    <Button icon
                      labelPosition='right'
                      color="green"
                      onClick={() => this.updateTask(item._id)}>
                      <Icon name ="check"/>
                      <span>Done</span>
                    </Button>
                    <Button icon
                      labelPosition='right'
                      color="yellow"
                      onClick={() => this.undoTask(item._id)}>
                      <Icon name ="undo"/>
                      <span>Undo</span>
                    </Button>

                    <Button icon
                      labelPosition='right'
                      color="red"
                      onClick={() => this.deleteTask(item._id)}>
                      <Icon name ="delete"/>
                      <span>Delete</span>
                    </Button>
                  </Card.Content>
                </Card.Content>
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
      // console.log(res);
    });
  };

  updateTask = id => {
    axios
      .put(endpoint + "/api/taskComplete/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        this.getTask();
      });
  };

  undoTask = id => {
    axios
      .put(endpoint + "/api/undoTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        this.getTask();
      });
  };

  deleteTask = id => {
    axios
      .delete(endpoint + "/api/deleteTask/" + id, {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        this.getTask();
      });
  };

  deleteAll = () => {
    axios
      .delete(endpoint + "/api/deleteAllTask", {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded"
        }
      })
      .then(res => {
        this.getTask();
      });
  };

  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h1" >
            TO DO LIST
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="task"
              onChange={this.onChange}
              value={this.state.task}
              fluid
              placeholder="Create Task"
            />
          </Form>
        </div>
        <div className="row"> 
          <Button 
            floated="right"
            color="red"
            onClick={() => this.deleteAll()}>
              Clear All
          </Button>
        </div>
        <div className="row">
          <Card.Group>{this.state.items}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;
