<template>
          <footer class="index-footer">
            <div class="d-flex flex-row">
                <img class="footer-logo" src="@/assets/bookflix.png" alt="Logo Bookflix">
            </div>
        </footer>
</template>

<script>
export default {
    name: 'Footer'
}
</script>

<style>
.index-footer {
    background-color: #232323; 
    color: white; 
    font-weight: bold;
}

.footer-logo {
    max-width: 8%;
    height: auto;
    margin: 3rem;
}
</style>