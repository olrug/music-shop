export default {
    name: 'App',
    data() {
      return {
        title: 'Первый коммит',
        description: 'Получаем данные...',
        counter: 0,
        users: [],
      }
    },
    methods: {
      incrementCounter() {
        this.counter += 1
      },
      GetUsers() {
        fetch('http://127.0.0.1:8090/api/users', {
          mode: 'cors',
          method: "get"
        })
        .then((response) => {
          if (!response.ok) {
            console.error('Ошибка при получении данных:', response.status, response.statusText);
            throw new Error(`Ошибка ${response.status}: ${response.statusText}`);
          }
          return response.json();
        })
        .then((data) => {
          console.log('Данные получены успешно:', data);
          this.users = data;
        })
        .catch((error) => {
          console.error('Ошибка при получении данных:', error);
        });
      }
    }
  }