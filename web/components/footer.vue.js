export default {
  props: {
    state: Object,
    deleteAllFunc: Function,
  },
  data() {
    return {
      showDanger: false,
      currentTheme: localStorage.getItem('theme') || 'auto',
    }
  },
  methods: {
    setTheme(theme) {
      this.currentTheme = theme
      localStorage.setItem('theme', theme)
      document.documentElement.setAttribute('data-theme', theme)
    },
    confirmDeleteAll() {
      if (confirm("Are you sure you want to delete all messages? This action cannot be undone.")) {
        this.deleteAllFunc()
        this.showDanger = false
      }
    },
  },
  beforeMount() {
    document.documentElement.setAttribute('data-theme', this.currentTheme)
  },
  template: `
    <footer class="footer is-fixed-bottom">
      <div class="level is-mobile">
        <div class="level-left">
          <div class="level-item" v-if="showDanger">
            <div class="field has-addons">
              <p class="control">
                <button class="button is-small" :class="{'is-link': currentTheme === 'auto'}" @click="setTheme('auto')" title="System Theme">
                  <span class="icon is-small"><i class="fas fa-magic"></i></span>
                  <span>Auto</span>
                </button>
              </p>
              <p class="control">
                <button class="button is-small" :class="{'is-link': currentTheme === 'light'}" @click="setTheme('light')" title="Light Mode">
                  <span class="icon is-small"><i class="fas fa-sun"></i></span>
                  <span>Light</span>
                </button>
              </p>
              <p class="control">
                <button class="button is-small" :class="{'is-link': currentTheme === 'dark'}" @click="setTheme('dark')" title="Dark Mode">
                  <span class="icon is-small"><i class="fas fa-moon"></i></span>
                  <span>Dark</span>
                </button>
              </p>
            </div>
          </div>
        </div>
        
        <div class="level-right">
          <div class="level-item">
            <div class="buttons">
              <a href="https://github.com/rapatao/sendgrid-mock" target="_blank" class="button is-light is-small">
                <span class="icon is-small"><i class="fab fa-github"></i></span>
                <span>GitHub</span>
              </a>
              
              <button class="button is-small is-white has-text-grey" @click="showDanger = !showDanger" title="Settings">
                <span class="icon is-small"><i class="fas fa-cog"></i></span>
              </button>

              <button v-if="showDanger" class="button is-danger is-outlined is-small ml-2" @click="confirmDeleteAll">
                <span class="icon is-small"><i class="fas fa-trash-alt"></i></span>
                <span>Delete all</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </footer>
  `
}
