<script setup lang="ts">
import { ref } from 'vue'
import { createEmptyBook, type Book } from '@/models/books'
import BackToHomeButton from '@/components/BackToHomeButton.vue'
import axios from 'axios'
import { useToast } from 'vue-toastification'
import router from '@/router'

const toast = useToast()

const addBook = ref<Book>(createEmptyBook())

async function handleSubmit() {
  try {
    const resp = await axios.post(`/api/books`, addBook.value)
    toast.success('Book added')
    router.push(`/books/${resp.data.id}`)
  } catch (error) {
    console.error('error adding book', error)
    toast.error('Book not added')
  }
}
</script>

<template>
  <BackToHomeButton />
  <div class="flex justify-center">
    <div class="w-2/3">
      <section class="container mx-auto px-4 py-8">
        <form @submit.prevent="handleSubmit">
          <div class="bg-white rounded-lg shadow-md p-6">
            <main>
              <h1 class="text-3xl font-bold mb-6">Add Book</h1>
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
                <div>
                  <label class="font-semibold block mb-1"
                    ><span class="text-red-500">*</span> Title:</label
                  >
                  <input
                    v-model="addBook.title"
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
                    v-model="addBook.author"
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
                    v-model="addBook.genre"
                    type="text"
                    class="w-full border rounded px-3 py-2"
                    required
                  />
                </div>
                <div>
                  <label class="font-semibold block mb-1">ISBN:</label>
                  <input
                    v-model="addBook.isbn"
                    type="text"
                    class="w-full border rounded px-3 py-2"
                  />
                </div>
                <div>
                  <label class="font-semibold block mb-1">Published Date:</label>
                  <input
                    v-model="addBook.published"
                    type="text"
                    class="w-full border rounded px-3 py-2"
                  />
                </div>
                <div>
                  <label class="font-semibold block mb-1">Publisher:</label>
                  <input
                    v-model="addBook.publisher"
                    type="text"
                    class="w-full border rounded px-3 py-2"
                  />
                </div>
              </div>
              <div class="mt-6">
                <label class="font-semibold block mb-1">Description:</label>
                <textarea
                  v-model="addBook.description"
                  class="w-full border rounded px-3 py-2"
                  rows="5"
                ></textarea>
              </div>
            </main>

            <button
              type="submit"
              class="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full w-full focus:outline-none focus:shadow-outline mb-4 mt-6"
            >
              Add Book
            </button>
          </div>
        </form>
      </section>
    </div>
  </div>
</template>
