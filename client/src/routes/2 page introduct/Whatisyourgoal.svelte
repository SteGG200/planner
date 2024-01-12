<script>
  import { onMount } from 'svelte';

  let inputArray = [];
  let inputValue = '';

  const handleFormSubmission = () => {
    inputArray.push(inputValue);
    console.log('Input Array:', inputArray);

    window.location.href = 'new-page.html';

    inputValue = '';
  };

  onMount(() => {
    const inputForm = document.getElementById('inputForm');
    const inputField = document.getElementById('inputField');
    const continueButton = document.getElementById('storeInput');

    inputForm.addEventListener('submit', (event) => {
      event.preventDefault();
      handleFormSubmission();
    });

    continueButton.addEventListener('click', handleFormSubmission);

    inputField.addEventListener('keypress', (event) => {
      if (event.keyCode === 13) {
        handleFormSubmission();
      }
    });
  });
</script>

<style>
  .bg-gray-800 {
    background-color: #1a202c;
  }

  .text-white {
    color: #ffffff;
  }

  .container {
    width: 100%;
    padding-right: 2rem;
    padding-left: 2rem;
    margin-right: auto;
    margin-left: auto;
  }

  .flex {
    display: flex;
  }

  .items-center {
    align-items: center;
  }

  .justify-center {
    justify-content: center;
  }

  .h-screen {
    height: 100vh;
  }

  .bg-gray-700 {
    background-color: #2d3748;
  }

  .p-8 {
    padding: 2rem;
  }

  .rounded-lg {
    border-radius: 0.5rem;
  }

  .shadow-lg {
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  }

  .w-96 {
    width: 24rem;
  }

  .text-2xl {
    font-size: 1.5rem;
    line-height: 2rem;
  }

  .font-semibold {
    font-weight: 600;
  }

  .mb-4 {
    margin-bottom: 1rem;
  }

  .text-center {
    text-align: center;
  }

  .mt-4 {
    margin-top: 1rem;
  }

  .block {
    display: block;
  }

  .text-gray-200 {
    color: #edf2f7;
  }

  .mt-1 {
    margin-top: 0.25rem;
  }

  .p-2 {
    padding: 0.5rem;
  }

  .border {
    border-width: 1px;
    border-color: #4a5568;
  }

  .rounded-md {
    border-radius: 0.375rem;
  }

  .w-full {
    width: 100%;
  }

  .hover\:bg-black:hover {
    background-color: #1a202c;
  }

  .bg-gray-800 {
    background-color: #2d3748;
  }

  .text-gray-800 {
    color: #2d3748;
  }
</style>

<div class="bg-gray-800 text-white">
  <div class="container mx-auto flex items-center justify-center h-screen">
    <div class="bg-gray-700 p-8 rounded-lg shadow-lg w-96">
      <h1 class="text-2xl font-semibold mb-4">PLANNER want to know: </h1>
      <form on:submit|preventDefault={handleFormSubmission} class="text-center" id="inputForm">
        <div class="mt-4">
          <label for="inputField" class="block text-gray-200">What is your goal?</label>
          <input bind:value={inputValue} type="text" id="inputField" class="mt-1 p-2 border rounded-md w-full text-gray-800" placeholder="Type something..." required />
        </div>
        <button type="submit" id="storeInput" class="mt-4 bg-gray-800 text-white p-2 rounded-md w-full hover:bg-black">Continue</button>
      </form>
    </div>
  </div>
</div>
