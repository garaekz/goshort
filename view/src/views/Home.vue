<template>
  <div class="home pt-36">
    <div class="flex items-center justify-center logo-wrapper">
      <img class="h-32" alt="GoShort Logo" src="../assets/goshort.svg">
    </div>
    <div class="mt-12 flex items-center justify-center bg-gray-600 bg-opacity-20 pt-6 pb-6">
      <Shortener/>
    </div>
    <div class="w-full flex justify-center pt-2">
      Previous <span class="pl-2 uppercase font-extrabold">Shorts</span>
    </div>
    <div class="w-full flex justify-center pb-8">
      <svg class="h-4" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" viewBox="0 0 350 350" xml:space="preserve">
      <g id="icon" style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: none; fill-rule: nonzero; opacity: 1;" transform="translate(-1.9444444444444287 -1.9444444444444287) scale(3.89 3.89)" >
        <polygon points="0,38.92 2.83,36.08 45,78.25 87.17,36.08 90,38.92 45,83.92 " style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(0,0,0); fill-rule: nonzero; opacity: 1;" transform="  matrix(1 0 0 1 0 0) "/>
        <polygon points="0,8.92 2.83,6.08 45,48.25 87.17,6.08 90,8.92 45,53.92 " style="stroke: none; stroke-width: 1; stroke-dasharray: none; stroke-linecap: butt; stroke-linejoin: miter; stroke-miterlimit: 10; fill: rgb(0,0,0); fill-rule: nonzero; opacity: 1;" transform="  matrix(1 0 0 1 0 0) "/>
      </g>
      </svg>
    </div>
    <div class="w-full flex justify-center">
      <ShortedList/>
    </div>
    <button @click="scrollToTop" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded fixed bottom-1 right-2 m-4">
      <svg
        class="h-4 fill-current text-white"
        xmlns:dc="http://purl.org/dc/elements/1.1/"
        xmlns:cc="http://creativecommons.org/ns#"
        xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
        xmlns:svg="http://www.w3.org/2000/svg"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
        xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
        viewBox="0 -256 1792 1792"
        id="svg3025"
        version="1.1"
        inkscape:version="0.48.3.1 r9886"
        sodipodi:docname="arrow_up_font_awesome.svg">
        <metadata
          id="metadata3035">
          <rdf:RDF>
            <cc:Work
              rdf:about="">
              <dc:format>image/svg+xml</dc:format>
              <dc:type
                rdf:resource="http://purl.org/dc/dcmitype/StillImage" />
            </cc:Work>
          </rdf:RDF>
        </metadata>
        <defs
          id="defs3033" />
        <sodipodi:namedview
          pagecolor="#ffffff"
          bordercolor="#666666"
          borderopacity="1"
          objecttolerance="10"
          gridtolerance="10"
          guidetolerance="10"
          inkscape:pageopacity="0"
          inkscape:pageshadow="2"
          inkscape:window-width="640"
          inkscape:window-height="480"
          id="namedview3031"
          showgrid="false"
          inkscape:zoom="0.13169643"
          inkscape:cx="896"
          inkscape:cy="896"
          inkscape:window-x="0"
          inkscape:window-y="25"
          inkscape:window-maximized="0"
          inkscape:current-layer="svg3025" />
        <g
          transform="matrix(1,0,0,-1,53.152542,1262.6441)"
          id="g3027">
          <path
            d="m 1611,565 q 0,-51 -37,-90 l -75,-75 q -38,-38 -91,-38 -54,0 -90,38 L 1024,693 V -11 q 0,-52 -37.5,-84.5 Q 949,-128 896,-128 H 768 q -53,0 -90.5,32.5 Q 640,-63 640,-11 V 693 L 346,400 q -36,-38 -90,-38 -54,0 -90,38 l -75,75 q -38,38 -38,90 0,53 38,91 l 651,651 q 35,37 90,37 54,0 91,-37 l 651,-651 q 37,-39 37,-91 z"
            id="path3029"
            inkscape:connector-curvature="0"
            style="fill:currentColor" />
        </g>
      </svg>
    </button>
    <div class="py-48">
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Shortener from '@/components/Shortener.vue';
import ShortedList from '@/components/ShortedList.vue';
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'Home',
  components: {
    Shortener,
    ShortedList,
  },
  computed: {
    ...mapGetters(['shorteds']),
  },
  mounted() {
  },
  methods: {
    ...mapActions({
      getSavedShorteds: 'GET_SAVED_SHORTEDS',
    }),
    scrollToTop() {
      window.scrollTo({
        top: 0,
        left: 0,
        behavior: 'smooth',
      });
    },
  },
};
</script>
<style scoped lang="scss">
.home {
  background: url('../assets/goshortbg.svg') no-repeat center center / cover;
}
</style>
