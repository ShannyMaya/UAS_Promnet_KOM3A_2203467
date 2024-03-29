import React, { Component } from 'react'
import UserService from '../services/UserService'

class ListUserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                users: []
        }
        this.addUser = this.addUser.bind(this);
        this.editUser = this.editUser.bind(this);
        this.deleteUser = this.deleteUser.bind(this);
    }

    deleteUserWithConfirmation(id) {
        const isConfirmed = window.confirm("Are you sure you want to delete this item?");
        
        if (isConfirmed) {
          // If the user confirms, proceed with the deletion
          this.deleteUser(id);
        }
        // If the user cancels, do nothing
      };

    deleteUser(id){
        UserService.deleteUser(id).then( res => {
            console.log("User deleted successfully");
            this.setState({users: 
                this.state.users.
                filter(user => user.id !== id)});
        });
    }
    viewUser(id){
        this.props.history.push(`/view-user/${id}`);
    }
    editUser(id){
        this.props.history.push(`/add-user/${id}`);
    }

    componentDidMount(){
        UserService.getUsers().then((res) => {
            if(res.data==null)
            {
                this.props.history.push('/add-user/_add');
            }
            this.setState({ users: res.data});
        });
    }

    addUser(){
        this.props.history.push('/add-user/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">
                     Items List</h2>
                 <div className = "row">
                    <button className="btn btn-primary"
                     onClick={this.addUser}> Add Item</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className 
                        = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th>Nama Barang</th>
                                    <th>Jumlah</th>
                                    <th>Harga Satuan</th>
                                    <th>Lokasi</th>
                                    <th>Deskripsi</th>
                                    <th>Aksi</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.users.map(
                                        user => 
                                        <tr key = {user.id}>
                                            <td>
                                                {user.nama_barang}
                                            </td>
                                            <td>
                                                {user.jumlah}
                                            </td>
                                            <td>
                                                {user.harga_satuan}
                                            </td>
                                            <td>
                                                {user.lokasi}
                                            </td>
                                            <td>
                                                {user.deskripsi}
                                            </td>
                                             <td>
                      <button onClick={ () => 
                          this.editUser(user.id)} 
                               className="btn btn-info">Update 
                                 </button>
                       <button style={{marginLeft: "10px"}}
                          //onClick={ () => this.deleteUser(user.id)} 
                          onClick={() => this.deleteUserWithConfirmation(user.id)}
                             className="btn btn-danger">Delete 
                                 </button>
                       <button style={{marginLeft: "10px"}} 
                           onClick={ () => this.viewUser(user.id)}
                              className="btn btn-info">Details 
                                  </button>
                                    </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>
                 </div>
            </div>
        )
    }
}

export default ListUserComponent
            </div>
        )
    }
}

export default ListUserComponent
