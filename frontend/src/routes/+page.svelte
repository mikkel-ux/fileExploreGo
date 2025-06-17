<script lang="ts">
	import { Greet } from '../lib/wailsjs/go/main/App';
	import { SayHi } from '../lib/wailsjs/go/main/Utils';
	import { GetFiles, GetDefaultDirs } from '../lib/wailsjs/go/goFiles/Files';

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

	interface FileDataMap {
		home: string;
		desktop: string;
		documents?: string;
		downloads?: string;
		pictures?: string;
		music?: string;
		videos?: string;
		[key: string]: string | undefined;
	}

	function handleClick() {
		GetDefaultDirs()
			.then((data: FileDataMap) => {
				console.log(data.home);
			})
			.catch((error: Error) => {
				console.log('Error in Test:', error);
			});
	}
</script>

<button class="btn" on:click={handleClick}>Say Hi</button>
