export default {
  props: {
    state: Object,
  },
  template: `
    <footer class="footer">
      <div class="content has-text-centered">
        <p>Search contains <strong>{{ state.total }}</strong> message(s).</p>
      </div>
    </footer>
  `
}
