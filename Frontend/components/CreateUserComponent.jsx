import React, { Component } from "react";
import UserService from "../services/UserService";

class CreateUserComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      // step 2
      id: this.props.match.params.id,
      nama_barang: "",
      jumlah: "",
      harga_satuan: "",
      lokasi: "",
      deskripsi: "",
    };
    this.changeNamaBarang = this.changeNamaBarang.bind(this);
    this.changeJumlah = this.changeJumlah.bind(this);
    // this.incrementJumlah = this.incrementJumlah.bind(this);
    // this.decrementJumlah = this.decrementJumlah.bind(this);
    this.changeHargaSatuan = this.changeHargaSatuan.bind(this);
    this.changeLokasi = this.changeLokasi.bind(this);
    this.changeDeskripsi = this.changeDeskripsi.bind(this);
    this.saveOrUpdateUser = this.saveOrUpdateUser.bind(this);
  }

  // step 3
  componentDidMount() {
    // step 4
    if (this.state.id === "_add") {
      return;
    } else {
      UserService.getUserById(this.state.id).then((res) => {
        let user = res.data;
        this.setState({
            nama_barang: user.nama_barang,
            jumlah: user.jumlah,
            harga_satuan: user.harga_satuan,
            lokasi: user.lokasi,
            deskripsi: user.deskripsi,
        });
      });
    }
  }
  saveOrUpdateUser = (e) => {
    e.preventDefault();
    let user = {
        nama_barang: this.state.nama_barang,
        jumlah: this.state.jumlah,
        harga_satuan: this.state.harga_satuan,
        lokasi: this.state.lokasi,
        deskripsi: this.state.deskripsi,
    };
    console.log("user => " + JSON.stringify(user));

    // step 5
    if (this.state.id === "_add") {
      UserService.createUser(user).then((res) => {
        this.props.history.push("/users");
      });
    } else {
      UserService.updateUser(user, this.state.id).then((res) => {
        this.props.history.push("/users");
      });
    }
  };

  changeNamaBarang = (event) => {
    this.setState({ nama_barang: event.target.value });
  };

  changeJumlah = (event) => {
    this.setState({ jumlah: event.target.value });
  };
  //incrementJumlah = () => {
    //this.setState((prevState) => ({
      //jumlah: prevState.jumlah + 1,
    //}));
  //};
  //decrementJumlah = () => {
    //this.setState((prevState) => ({
      //jumlah: Math.max(1, prevState.jumlah - 1),
    //}));
  //};
  changeHargaSatuan = (event) => {
    this.setState({ harga_satuan: event.target.value });
  };

  changeLokasi = (event) => {
    console.log(event.target.value);
    this.setState({ lokasi: event.target.value });
  };

  changeDeskripsi = (event) => {
    this.setState({ deskripsi: event.target.value });
  };

  cancel() {
    this.props.history.push("/users");
  }

  getTitle() {
    if (this.state.id === "_add") {
      return <h3 className="text-center">Add Item</h3>;
    } else {
      return <h3 className="text-center">Update Item</h3>;
    }
  }
  render() {
    return (
      <div>
        <br></br>
        <div className="container">
          <div className="row">
            <div className="card col-md-6 offset-md-3 offset-md-3">
              {this.getTitle()}
              <div className="card-body">
                <form>
                  <div className="form-group">
                    <label> Nama Barang: </label>
                    <input
                      placeholder="Nama Barang"
                      name="nama_barang"
                      className="form-control"
                      value={this.state.nama_barang}
                      onChange={this.changeNamaBarang}
                    />
                  </div>
                  <div className="form-group">
                    <label> Jumlah: </label>
                    <input
                      placeholder="Jumlah"
                      name="jumlah"
                      className="form-control"
                      value={this.state.jumlah}
                      onChange={this.changeJumlah}
                      
                    />
                    {/* <div className="input-group-prepend">
                    <button
                      type="button"
                      className="btn btn-outline-secondary"
                      onClick={this.decrementJumlah}
                    >
                      -
                    </button>
                    <div className="input-group-append">
                    <button
                      type="button"
                      className="btn btn-outline-secondary"
                      onClick={this.incrementJumlah}
                    >
                      +
                    </button>
            </div>
            </div>
                  </div>
                  <div className="wrapper">
                      <span className="minus" onClick={this.decrementJumlah}>
                        -
                      </span>
                      <span className="num" name="jumlah">
                        {this.state.jumlah}
                      </span>
                     <span className="plus" onClick={this.incrementJumlah}>
                        +
                      </span> */}
                  </div>
                  <div className="form-group">
                    <label> Harga Satuan: </label>
                    <input
                      placeholder="Harga Satuan"
                      name="harga_satuan"
                      className="form-control"
                      value={this.state.harga_satuan}
                      onChange={this.changeHargaSatuan}
                    />
                  </div>
                  <div className="form-group">
                    <label> Lokasi: </label>
                    <select
                      name="lokasi"
                      className="form-control"
                      value={this.state.lokasi}
                      onChange={this.changeLokasi}
                    >
                      <option value="Bandung">Bandung</option>
                      <option value="Jakarta">Jakarta</option>
                      <option value="Denpasar">Denpasar</option>
                      <option value="Manokwari">Manokwari</option>
                    </select>
                  </div>
                  <div className="form-group">
                    <label> Deskripsi: </label>
                    <input
                      placeholder="Deskripsi"
                      name="deskripsi"
                      className="form-control"
                      value={this.state.deskripsi}
                      onChange={this.changeDeskripsi}
                    />
                  </div>

                  <button
                    className="btn btn-success"
                    onClick={this.saveOrUpdateUser}
                  >
                    Save
                  </button>
                  <button
                    className="btn btn-danger"
                    onClick={this.cancel.bind(this)}
                    style={{ marginLeft: "10px" }}
                  >
                    Cancel
                  </button>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default CreateUserComponent;
