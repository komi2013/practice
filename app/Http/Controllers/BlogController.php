<?php
namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Storage;
use Illuminate\Support\Facades\DB;

class BlogController extends Controller {

    public function index(Request $request) {
$user = DB::table('test1')->first();
echo $user->test_txt;
die('programming practice');
    }
}