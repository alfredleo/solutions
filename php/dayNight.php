<?php
declare(strict_types=1);

function getTime(): string
{
    $time = 'день';
    /* Здесь начинается ночная магия... */

    /* ...а здесь она заканчивается */
    var_dump("Сейчас $time"); // string(21) "Сейчас день"
    return $time;
}

$time = getTime();
var_dump("Сейчас $time"); //string(21) "Сейчас ночь"