    <div class="description">
		<p><<<.Index>>></p>
		<<<range $idx, $site := .Sites >>>
		<div>
			<span><<<$site.Name>>></span>: 
			<span> <<<$site.Url>>></span>
		</div>
		<<< end >>>
    </div>