<template>
          <header id="topbar" class="navbar navbar-expand navbar-dark topbar mb-4 fixed-top shadow">
            <router-link to="/" class="sidebar-brand d-flex align-items-center justify-content-center">
                <img class="sidebar-logo" src="@/assets/bookflix.png" alt="Logo Bookflix">
            </router-link>
           
            <form class="d-none d-sm-inline-block form-inline mr-auto ml-auto my-2 my-md-0 mw-100 navbar-search">
                <div class="input-group">
                    <input type="text" class="form-control bg-dark border-0" placeholder="Search for book, magazine..."
                        aria-label="Search" aria-describedby="basic-addon2">
                    <div class="input-group-append">
                        <button class="btn btn-dark" type="button">
                            <em class="fa fa-search fa-sm"></em>
                        </button>
                    </div>
                </div>
            </form>

            <ul class="navbar-nav ml-auto">
                <li class="nav-item dropdown no-arrow d-sm-none show">
                    <a class="nav-link dropdown" href="#" id="searchDropdown" role="button" data-toggle="dropdown"
                        aria-haspopup="true" aria-expanded="true">
                        <em class="fa fa-search fa-fw"></em>
                    </a>
                    <!-- Dropdown - Messages -->
                    <div class="dropdown-menu dropdown-menu-right p-3 shadow" aria-labelledby="searchDropdown">
                        <form class="form-inline mr-auto w-100 navbar-search">
                            <div class="input-group">
                                <input type="text" class="form-control bg-dark border-0"
                                    placeholder="Search for book, magazine ..." aria-label="Search"
                                    aria-describedby="basic-addon2">
                                <div class="input-group-append">
                                    <button class="btn btn-dark" type="button">
                                        <em class="fa fa-search fa-sm"></em>
                                    </button>
                                </div>
                            </div>
                        </form>
                    </div>
                </li>
                <!-- Nav Item - User Information -->
                <li class="nav-item dropdown no-arrow">
                    <a class="nav-link dropdown-toggle" href="#" id="userDropdown" role="button" data-toggle="dropdown"
                        aria-haspopup="true" aria-expanded="false">
                        <img class="img-profile rounded-circle" src="@/assets/userImage.jpg" alt="Foto do Usuário">
                    </a>
                    <!-- Dropdown - User Information -->
                    <div class="dropdown-menu dropdown-menu-right shadow animated--grow-in"
                        aria-labelledby="userDropdown">
                        <a class="dropdown-item" href="#">
                            <em class="fa fa-user fa-sm fa-fw mr-2 text-gray-400"></em>
                            Profile
                        </a>
                        <a class="dropdown-item" href="#">
                            <em class="fa fa-cogs fa-sm fa-fw mr-2 text-gray-400"></em>
                            Settings
                        </a>
                        <a class="dropdown-item" href="#">
                            <em class="fa fa-list fa-sm fa-fw mr-2 text-gray-400"></em>
                            Activity Log
                        </a>
                        <div class="dropdown-divider"></div>
                        <a class="dropdown-item" href="#" data-toggle="modal" data-target="#logoutModal">
                            <em class="fa fa-sign-out-alt fa-sm fa-fw mr-2 text-gray-400"></em>
                            Logout
                        </a>
                    </div>
                </li>

            </ul>
          </header>
</template>

<script>
export default {
    name: 'Topbar'
}
</script>

<style>
.topbar {
    background-color: #000000;
}

.sidebar-brand {
    height: 4.375rem;
    text-decoration: none;
    font-size: 1rem;
    font-weight: 800;
    padding: 1.5rem 1rem;
    text-align: center;
    z-index: 1;
}

.sidebar-logo {
    max-width: 50%;
    height: auto;
}

.img-profile {
    height: 2rem;
    width: 2rem;
}

.topbar .navbar-search {
    width: 50rem;
    color: white;
}

.topbar .dropdown .dropdown-menu {
    width: calc(100% - 1.5rem);
    right: .75rem;
}

.input-group>.custom-select:not(:last-child), .input-group>.form-control:not(:last-child) {
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
    color: white;
    background-color: #1c1c1c;
}
</style>