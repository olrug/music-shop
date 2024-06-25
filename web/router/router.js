import { createRouter, createWebHistory } from 'vue-router';
import HomePage from '../pages/HomePage.vue';
import CatalogPage from '../pages/CatalogPage.vue'
import Navbar from '../components/Navbar.vue';

const routes = [
  {
    path: '/',
    name: '/home',
    components: {
      default: HomePage,
      navbar: Navbar,
    },
    meta: { title: 'Главная страница' }
  },
  {
    path: '/catalog',
    name: '/catalog',
    components: {
      default: CatalogPage,
      navbar: Navbar,
    },
    meta: { title: 'Каталог товаров' }
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});


export default router;