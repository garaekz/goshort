<template>
  <div class="w-full flex justify-center">
    <div class="w-full md:w-44r">
      <div class="shortened-res px-2">
        <h1 v-if="shorted" v-clipboard:copy="shorted" @click="copyToClipboard()" class="flex justify-center items-center m-1 font-medium text-5xl">
          <transition name="fade" v-if="!copying">
            <!-- icon here -->
            <font-awesome-icon class="copy-icon mr-2" :icon="['far', 'copy']" />
          </transition>
          <transition name="fade" v-else>
            <!-- icon here -->
            <font-awesome-icon class="check-icon mr-2" icon="check" />
          </transition>
          {{ shorted }}
        </h1>
      </div>
      <div :class="[error ? 'pb-4' : 'pb-12', shorted ? 'pt-4' : 'pt-12']" class="flex flex-wrap items-stretch relative md:w-44r w-full px-2">
        <input
          type="text"
          class="search-input flex-shrink flex-grow flex-auto leading-normal w-px flex-1 border h-12 border-grey-light rounded px-3 relative focus:outline-none focus:ring-1"
          placeholder="Paste your URL and GOSHORT"
          v-model="url"
          @keyup.enter="onSaveURL"
        />
        <div class="flex -mr-px">
          <button @click.prevent="onSaveURL" class="flex items-center bg-urbano-green leading-normal rounded rounded-l-none border border-l-0 border-urbano-green px-3 whitespace-no-wrap text-white h-12 text-sm focus:outline-none hover:bg-gray-500 hover:border-gray-500">
            <svg class="h-4 fill-current text-white" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M476 3.2L12.5 270.6c-18.1 10.4-15.8 35.6 2.2 43.2L121 358.4l287.3-253.2c5.5-4.9 13.3 2.6 8.6 8.3L176 407v80.5c0 23.6 28.5 32.9 42.5 15.8L282 426l124.6 52.2c14.2 6 30.4-2.9 33-18.2l72-432C515 7.8 493.3-6.8 476 3.2z"/></svg>
          </button>
        </div>
      </div>
      <div v-if="error" class="flex justify-center items-center m-1 font-medium py-1 px-2 bg-white rounded-md text-red-100 bg-red-700 border border-red-700 ">
        <div slot="avatar">
          <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-alert-octagon w-5 h-5 mx-2">
            <polygon points="7.86 2 16.14 2 22 7.86 22 16.14 16.14 22 7.86 22 2 16.14 2 7.86 7.86 2"></polygon>
            <line x1="12" y1="8" x2="12" y2="12"></line>
            <line x1="12" y1="16" x2="12.01" y2="16"></line>
          </svg>
        </div>
        <div class="text-xl font-normal  max-w-full flex-initial">
          There is an error in your code</div>
        <div class="flex flex-auto flex-row-reverse">
          <div @click="error = null">
            <svg xmlns="http://www.w3.org/2000/svg" width="100%" height="100%" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-x cursor-pointer hover:text-red-400 rounded-full w-5 h-5 ml-2">
              <line x1="18" y1="6" x2="6" y2="18"></line>
              <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
          </div>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import mixin from '@/mixins';

export default {
  name: 'Shortener',
  mixins: [mixin],
  data: () => ({
    url: '',
    shorted: null,
    error: null,
    copying: false,
  }),
  computed: {
    ...mapGetters(['shorteds']),
  },
  methods: {
    ...mapActions({
      saveURL: 'SAVE_URL',
    }),
    onSaveURL() {
      this.saveURL(this.url).then((res) => {
        this.shorted = this.urlify(res.data.code);
        this.url = null;
      }).catch((err) => {
        // Edit this if you need to validate multiple fields
        const { field, data } = err.data.details[0];
        this.error = `${field === 'url' ? field.toUpperCase() : field} ${data}`;
      });
    },
    copyToClipboard() {
      this.copying = true;
      setTimeout(() => {
        this.copying = false;
        return true;
      }, 1000);
    },
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
