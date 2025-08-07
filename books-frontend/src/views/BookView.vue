<script setup lang="ts">
import router from '@/router'
import BackToHomeButton from '@/components/BackToHomeButton.vue'
import axios from 'axios'
import { useRoute } from 'vue-router'
import { ref, onMounted } from 'vue'
import { useToast } from 'vue-toastification'
import { createEmptyBook, type Book } from '@/models/books'

const route = useRoute()
const toast = useToast()

const bookId = route.params.id

const book = ref<Book>(createEmptyBook())

async function deleteBook() {
  try {
    const confirm = window.confirm('Confirm deleting this book?')
    if (confirm) {
      await axios.delete(`/api/books/${bookId}`)
      toast.success('Book deleted')
      router.push('/')
    }
  } catch (error) {
    console.error('error deleting book', error)
    toast.error('Book not deleted')
  }
}

onMounted(async () => {
  try {
    const resp = await axios.get(`/api/books/${bookId}`)
    book.value = resp.data
  } catch (error) {
    console.error('error fetching book', error)
  }
})
</script>

<template>
  <BackToHomeButton />
  <div class="flex justify-center">
    <div class="w-2/3">
      <section class="container mx-auto px-4 py-8">
        <div class="bg-white rounded-lg shadow-md p-6">
          <main>
            <h1 class="text-3xl font-bold mb-6">{{ book.title }}</h1>
            <div class="m-3">
              <h3 class="font-semibold">Author:</h3>
              <p>{{ book.author }}</p>
            </div>
            <div class="m-3">
              <h3 class="font-semibold">Genre:</h3>
              <p>{{ book.genre }}</p>
            </div>
            <div class="m-3">
              <h3 class="font-semibold">ISBN:</h3>
              <p>{{ book.isbn }}</p>
            </div>
            <div class="m-3">
              <h3 class="font-semibold">Publish Data:</h3>
              <p>{{ book.published }}</p>
            </div>
            <div class="m-3">
              <h3 class="font-semibold">Publisher:</h3>
              <p>{{ book.publisher }}</p>
            </div>
            <div class="m-3">
              <h3 class="font-semibold">Description:</h3>
              <p>{{ book.description }}</p>
            </div>
          </main>

          <div class="bg-white p-6 rounded-lg shadow-md mt-6">
            <h3 class="text-xl font-bold mb-6">Manage Book</h3>
            <RouterLink
              :to="`/books/edit/${bookId}`"
              class="bg-green-500 hover:bg-green-600 text-white text-center font-bold py-2 px-4 rounded-full w-full focus:outline-none focus:shadow-outline mt-4 block"
              >Edit Book</RouterLink
            >
            <button
              @click="deleteBook"
              class="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded-full w-full focus:outline-none focus:shadow-outline mt-4 block"
            >
              Delete Book
            </button>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>
