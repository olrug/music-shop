<template>
  <NavbarComponent />
  <main class="content">
      <router-view />
  </main>
</template>

<script>
  import NavbarComponent from './components/Navbar.vue';

  export default {
      name: 'App',
      components: {
          NavbarComponent
      }
  }
</script>