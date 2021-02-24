export default {
  methods: {
    urlify: (code, protocol) => {
      const start = protocol ? `${window.location.protocol}//` : '';
      const response = `${start}${window.location.host}/${code}`;
      return response;
    },
  },
};
