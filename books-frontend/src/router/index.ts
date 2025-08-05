import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
// import BooksView from '@/views/BooksView.vue'
import BookView from '@/views/BookView.vue'
import EditBookView from '@/views/EditBookView.vue'
import AddBookView from '@/views/AddBookView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/books/:id',
      name: 'book',
      component: BookView,
    },
    {
      path: '/books/edit/:id',
      name: 'edit-book',
      component: EditBookView,
    },
    {
      path: '/books/add',
      name: 'add-book',
      component: AddBookView,
    },
  ],
})

export default router
