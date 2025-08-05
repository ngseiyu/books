<script setup lang="ts">
import { useRoute } from 'vue-router'
import { onMounted, ref } from 'vue'
import { createEmptyBook, type Book } from '@/models/books'
import BackToHomeButton from '@/components/BackToHomeButton.vue'
import axios from 'axios'
import { useToast } from 'vue-toastification'
import router from '@/router'

const route = useRoute()
const toast = useToast()

const bookId = parseInt(route.params.id as string, 10)

const book = ref<Book>(createEmptyBook())
const editBook = ref<Book>(createEmptyBook())

async function handleSubmit() {
  try {
    const resp = await axios.post(`/api/books/${bookId}`, editBook.value)
    toast.success('Book updated')
    router.push(`/books/${resp.data.ID}`)
  } catch (error) {
    console.error('error updating book', error)
    toast.error('Book not updated')
  }
}

function cancelEdit() {
  router.back()
}

onMounted(async () => {
  try {
    const resp = await axios.get(`/api/books/${bookId}`)
    book.value = resp.data
    editBook.value = {
      id: bookId,
      title: book.value.title,
      author: book.value.author,
      genre: book.value.genre,
      description: book.value.description,
      isbn: book.value.isbn,
      published: book.value.published,
      publisher: book.value.publisher,
    }
  } catch (error) {
    console.error('failed to fetch book', error)
  }
})
</script>

<template>
  <BackToHomeButton />
  <section class="container mx-auto px-4 py-8">
    <form @submit.prevent="handleSubmit">
      <div class="bg-white rounded-lg shadow-md p-6">
        <main>
          <h1 class="text-3xl font-bold mb-6">Edit Book</h1>
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <div>
              <label class="font-semibold block mb-1"
                ><span class="text-red-500">*</span> Title:</label
              >
              <input
                v-model="editBook.title"
                type="text"
                class="w-full border rounded px-3 py-2"
                required
              />
            </div>
            <div>
              <label class="font-semibold block mb-1"
                ><span class="text-red-500">*</span> Author:</label
              >
              <input
                v-model="editBook.author"
                type="text"
                class="w-full border rounded px-3 py-2"
                required
              />
            </div>
            <div>
              <label class="font-semibold block mb-1"
                ><span class="text-red-500">*</span> Genre:</label
              >
              <input
                v-model="editBook.genre"
                type="text"
                class="w-full border rounded px-3 py-2"
                required
              />
            </div>
            <div>
              <label class="font-semibold block mb-1">ISBN:</label>
              <input v-model="editBook.isbn" type="text" class="w-full border rounded px-3 py-2" />
            </div>
            <div>
              <label class="font-semibold block mb-1">Published Date:</label>
              <input
                v-model="editBook.published"
                type="text"
                class="w-full border rounded px-3 py-2"
              />
            </div>
            <div>
              <label class="font-semibold block mb-1">Publisher:</label>
              <input
                v-model="editBook.publisher"
                type="text"
                class="w-full border rounded px-3 py-2"
              />
            </div>
          </div>
          <div class="mt-6">
            <label class="font-semibold block mb-1">Description:</label>
            <textarea
              v-model="editBook.description"
              class="w-full border rounded px-3 py-2"
              rows="5"
            ></textarea>
          </div>
        </main>

        <!-- <div class="bg-white p-6 rounded-lg shadow-md mt-6">
          <h3 class="text-xl font-bold mb-4">Manage Book</h3> -->
        <button
          type="submit"
          class="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full w-full focus:outline-none focus:shadow-outline mb-4 mt-6"
        >
          Save Changes
        </button>
        <button
          type="button"
          @click="cancelEdit"
          class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-full w-full focus:outline-none focus:shadow-outline"
        >
          Cancel
        </button>
        <!-- </div> -->
      </div>
    </form>
  </section>
</template>
