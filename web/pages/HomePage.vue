<template>
    <div id="app" class="container">
      <header>
        <h1>{{ title }}</h1>
      </header>
      <main>
        <div>
          <p>{{ description }}</p>
          <button @click="GetUsers" class="btn">Нажми меня</button>
        </div>
        <div class="user-list">
          <div class="user-card" v-for="user in users" :key="user.id">
            <div class="user-id">{{ user.UserId }}</div>
            <div class="user-name">{{ user.Name }}</div>
            <div class="user-second-name">{{ user.SecondName }}</div>
            <div class="user-middle-name">{{ user.MiddleName }}</div>
            <div class="user-email">{{ user.Email }}</div>
            <div class="user-phone-number">{{ user.PhoneNumber }}</div>
          </div>
        </div>
      </main>
    </div>
  </template>
  
  <script src="../js/app.js"></script>
  
  <style>
  body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
    background-color: #f5f5f5;
  }
  
  .container {
    max-width: 800px;
    margin: 40px auto;
    padding: 20px;
    background-color: #fff;
    border: 1px solid #ddd;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }
  
  header {
    background-color: #333;
    color: #fff;
    padding: 10px;
    text-align: center;
  }
  
  h1 {
    margin: 0;
    font-size: 24px;
  }
  
  main {
    padding: 20px;
  }
  
  .btn {
    background-color: #4CAF50;
    color: #fff;
    border: none;
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
  }
  
  .btn:hover {
    background-color: #3e8e41;
  }
  
  .user-list {
    margin-top: 20px;
  }
  
  .user-card {
    background-color: #fff;
    padding: 10px;
    border: 1px solid #ddd;
    margin-bottom: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }
  
  .user-id,.user-name,.user-second-name,.user-middle-name,.user-email,.user-phone-number,.user-password {
    margin-bottom: 10px;
  }
  
  .user-id:before {
    content: "ID: ";
  }
  
  .user-name:before {
    content: "Имя: ";
  }
  
  .user-second-name:before {
    content: "Фамилия: ";
  }
  
  .user-middle-name:before {
    content: "Отчество: ";
  }
  
  .user-email:before {
    content: "Email: ";
  }
  
  .user-phone-number:before {
    content: "Телефон: ";
  }
  
  .user-password:before {
    content: "Пароль: ";
  }
  
  </style>