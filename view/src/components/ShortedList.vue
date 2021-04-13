<template>
  <div class="md:w-3/6 w-full grid md:grid-cols-2 md:gap-4 grid-cols-1">
    <div v-for="shorted in shorteds" :key="shorted.code">
      <article class="grid grid-cols-8 bg-white shadow-sm lg:max-w-2xl p-4 lg:col-span-2">
        <qr-code
          :text="urlify(shorted.code, true)"
          class="w-full col-span-3"
        ></qr-code>
        <div class="flex flex-col col-span-5">
          <div class="flex">
            <div class="w-full"></div>
            <div class="justify-self-end">
              <button @click.prevent="deleteURL(shorted.code)" class="flex items-center text-3xl focus:outline-none hover:text-gray-500">
                <font-awesome-icon class="text-gray-500 hover:text-gray-600" icon="times" />
              </button>
            </div>
          </div>
          <div class="w-full h-full pl-4 flex flex-col" style="overflow-wrap: break-word;">
            <h2 class="text-gray-800 text-xl font-bold">{{ urlify(shorted.code) }}</h2>
            {{ shorted.original_url }}
            <button v-clipboard:copy="urlify(shorted.code)" class="w-full bg-gray-200 mt-auto font-bold text-gray-500 hover:text-gray-600 py-2">
              <font-awesome-icon class="check-icon mr-2" icon="clipboard" /> Copy to clipboard
            </button>
          </div>
        </div>
      </article>
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import mixin from '@/mixins';

export default {
  name: 'ShortedList',
  mixins: [mixin],
  data: () => ({
  }),
  computed: {
    ...mapGetters(['shorteds']),
  },
  mounted() {
    this.getSavedShorteds();
  },
  methods: {
    ...mapActions({
      deleteURL: 'DELETE_SHORTED',
      getSavedShorteds: 'GET_SAVED_SHORTEDS',
    }),
  },
};
</script>

<style scoped lang="scss">
.search-input{
  font-size: 21px;
  &::placeholder {
    font-weight: 100;
    font-size: 28px;
  }
}
</style>
