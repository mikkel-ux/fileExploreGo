<script lang="ts">
	import { Greet } from '../lib/wailsjs/go/main/App';
	import { SayHi } from '../lib/wailsjs/go/main/Utils';
	import { GetFiles } from '../lib/wailsjs/go/goFiles/Files';

	type FileData = {
		name: string;
		path: string;
		size: string;
		extension: string;
		created: string;
		modified: string;
		accessed: string;
		fileType: string;
		isHidden: boolean;
		isReadOnly: boolean;
		base64?: string;
	};

	function handleClick() {
		GetFiles('C:/Users/rumbo/.testFoulderForFE')
			.then((data: FileData[]) => {
				console.log(data);
			})
			.catch((error: Error) => {
				console.log('Error in Test:', error);
			});
	}
</script>

<button class="btn" on:click={handleClick}>Say Hi</button>
