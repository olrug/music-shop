<template>
    <div class="Navbar">
      <nav>
        <ul>
          <li>
            <router-link active-class="active" to="/">Главная страница</router-link>
          </li>
          <li>
            <router-link active-class="active" to="/catalog">Каталог</router-link>
          </li>
        </ul>
      </nav>
    </div>
  </template>
  
  <style>
  .Navbar {
    background-color: #f8f8f8;
    padding: 10px;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    z-index: 999;
  
    display: flex;
    justify-content: center;
  }
  
  nav ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    display: flex;
  }
  
  nav ul li {
    margin-right: 10px;
  }
  
  nav ul li a {
    text-decoration: none;
    color: #333;
    padding: 5px;
    border-radius: 5px;
  }
  
  nav ul li a:hover {
    background-color: #333;
    color: #fff;
  }
  
  .active {
    background-color: #333;
    color: #fff;
  }
  </style>