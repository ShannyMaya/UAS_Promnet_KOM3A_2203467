import React, { Component } from "react";
import UserService from "../services/UserService";

class ViewUserComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      id: this.props.match.params.id,
      user: {},
    };
  }

  componentDidMount() {
    UserService.getUserById(this.state.id).then((res) => {
      this.setState({ user: res.data });
    });
  }

  render() {
    return (
      <div>
        <br></br>
        <div className="card col-md-6 offset-md-3">
          <h3 className="text-center">View Item Details</h3>
          <div className="card-body">
            <div className="row">
              <label> Nama Barang: </label>
              <div> {this.state.user.nama_barang}</div>
            </div>
            <div className="row">
              <label> Jumlah: </label>
              <div> {this.state.user.jumlah}</div>
            </div>
            <div className="row">
              <label> Harga Satuan: </label>
              <div> {this.state.user.harga_satuan}</div>
            </div>
            <div className="row">
              <label> Lokasi: </label>
              <div> {this.state.user.lokasi}</div>
            </div>
            <div className="row">
              <label> Deskripsi: </label>
              <div> {this.state.user.deskripsi}</div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default ViewUserComponent;
