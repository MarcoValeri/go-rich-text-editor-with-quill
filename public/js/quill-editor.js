window.onload = () => {
    // Initialize Quill editor
    const quill = new Quill('#editor', {
        theme: 'snow'
    });

    // DOM variables
    const form = document.getElementById('form');
    const formContentInput = document.getElementById('form-content-input');

    /**
     * JS submit event that gets data from the form
     * and
     * pass quill content to the hidden input
     */
    form.addEventListener('submit', () => {
        const getContent = quill.root.innerHTML;
        formContentInput.value = getContent;
    })
}