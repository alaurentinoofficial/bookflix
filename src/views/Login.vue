<template>
  <div class="body-wrapper">
                <div class="row justify-content-center mt-auto">
                    <img class="bookflix-logo" src="@/assets/bookflix.png" alt="Logo Bookflix">
                </div>
                <div class="row justify-content-center mt-4 mb-auto">
                    <div class="login-content col-3 p-3">
                        <h1 class="login-title">Login</h1>
                        <form class="login-form">
                            <div class="form-group">
                              <input type="email" class="form-control login-email-form" id="InputEmail" placeholder="Email" v-on:focusout="event => email = event.target.value">
                            </div>
                            <div class="form-group login-password-form-group">
                              <input type="password" class="form-control mb-1 login-password-input  " id="InputPassword" placeholder="Senha" v-on:focusout="event => password = event.target.value">
                              <a class="login-forgot-password" href="#">Esqueceu a senha?</a>
                            </div>
                            <button class="btn btn-primary w-100 mt-2 login-button" v-on:click="handleLogin()">Login</button>
                            <div class="text-center mt-4">
                                <span>Não tem cadastro?</span><a class="login-enroll" href="#"> Inscreva-se</a>
                            </div>
                          </form>
                    </div>
                </div>
            </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      email: "",
      password: "",
      loading: false,
      message: ''
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    }
  },
  created() {
    if (this.loggedIn) {
      this.$router.push('/profile');
    }
  },
  methods: {
    handleLogin() {
      if (this.email && this.password) {
        this.$store.dispatch('auth/login', { email: this.email, password: this.password }).then(
          () => {
            this.$router.push('/profile');
          },
          error => {
            this.message =
              (error.response && error.response.data) ||
              error.message ||
              error.toString();
          }
        );
      }
    }
  }
};
</script>

<style scoped>
#content-wrapper {
    background-color: #000000;
    height: 100vh;
    width: 100%;
    overflow-x: hidden;
}

.body-wrapper {
    width:50%;
    margin-top: auto;
    margin-bottom: auto;
    margin-left: auto;
    margin-right: auto;
}

.bookflix-logo {
    width: 15%;
    height: 100%;
}

.login-title {
    color: white;
    margin-top: 1.5rem;
    font-size: 1.5rem;
    text-align: center;
    font-weight: bolder;
}

hr .span-div-login { 
    display: inline-block;
    border-color: #FFFFFF;
    border: 1px inset;
    overflow: hidden;
}

.login-content {
    background-color: #232323; 
    color: white;
}

.login-form {
    margin-top: 4rem;
}

.login-email-form {
    height: 3.5rem;
}

.login-password-form-group {
    margin-top: 3rem;
}

.login-password-input {
    height: 3.5rem;
}

.login-forgot-password {
    color: #0090E4;
}

.login-button {
    background-color: #db2321; 
    border-color: #db2321; 
    color: white;
}

.login-enroll {
    color: #0090E4;
}
</style>