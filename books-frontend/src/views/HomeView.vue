<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import BookListings from '@/components/BookListings.vue'
import type { Book } from '@/models/books'

const books = ref<Book[]>([])

onMounted(async () => {
  try {
    const resp = await axios.get('/api/books')
    books.value = resp.data
  } catch (error) {
    console.error('error fetching books', error)
  }
})
</script>

<template>
  <section class="p-6">
    <div class="flex justify-center mb-4">
      <h1 class="text-3xl font-bold">NG Books</h1>
    </div>
    <div class="flex justify-center mb-6">
      <h2 class="text-xl font-semibold">Books Listing</h2>
    </div>
    <div class="flex justify-center">
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <RouterLink
          :to="'/books/add'"
          class="bg-white rounded-lg shadow-md hover:shadow-lg hover:bg-gray-50 transition-all duration-200 m-3 p-4 border border-gray-200 flex items-center justify-center"
        >
          <div class="text-center">
            <div class="text-5xl text-gray-400 mb-2">＋</div>
            <h3 class="text-lg font-semibold text-gray-700">Add New Book</h3>
          </div>
        </RouterLink>
        <BookListings v-for="book in books" :key="book.id" :book="book" />
      </div>
    </div>
  </section>
</template>
