<script>
  import { appDataDir, join, BaseDirectory } from "@tauri-apps/api/path";

  import { convertFileSrc } from "@tauri-apps/api/core";
  import { onMount } from "svelte";
  let imageUrl = "C://Users//rumbo//OneDrive//Billeder//myNewImage.png";

  const getImageUrl = async () => {
    return convertFileSrc(imageUrl);
  };

  onMount(async () => {
    let test = await getImageUrl();
    console.log("Test image URL:", test);
  });
</script>

<img
  src="http://asset.localhost/C%3A%2F%2FUsers%2F%2Frumbo%2F%2FOneDrive%2F%2FBilleder%2F%2FmyNewImage.png
"
  alt=""
/>
